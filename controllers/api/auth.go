package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controllers/mailer"
	"EvelyApi/controllers/parser"
	. "EvelyApi/middleware"
	"EvelyApi/models"
	. "EvelyApi/models/collections"
	. "EvelyApi/models/documents"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db *models.EvelyDB
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *models.EvelyDB) *AuthController {
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         db,
	}
}

// SendMail runs the send_mail action.
func (c *AuthController) SendMail(ctx *app.SendMailAuthContext) error {

	// メールアドレスが使用可能か検査
	email := ctx.Payload.Email
	if c.db.Users.Exists(Keys{"mail.email": email, "mail.state": STATE_OK}) {
		return ctx.BadRequest(goa.ErrBadRequest("\"" + email + "\" is already in use."))
	}

	// トークンを発行する
	claims := jwtgo.MapClaims{
		"scopes":     "api:access",
		"email":      email,
		"created_at": time.Now(),
	}
	token, err := NewToken(claims)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// メール送信
	url := "http://localhost:8888/verify_email?token=" + token
	err = mailer.SendSignUpConfirmMail(email, url)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// 認証待ちユーザーをDBに保存
	u := &UserModel{
		Mail: &Mail{
			Email: email,
			Token: token,
			State: STATE_PENDING,
		},
	}
	err = c.db.Users.Save(u, Keys{"mail.email": u.Mail.Email})
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}
	return ctx.OK([]byte("Success!!"))
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {

	// ユーザーが存在するか検索
	p := ctx.Payload
	user, err := c.db.Users.FindOne(Keys{"id": p.ID})
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest("The ID and password you entered did not match."))
	}

	// IDとパスワードが一致するかを検査
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest("The ID and password you entered did not match."))
	}

	// JWTを生成して返す
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
	}
	token, err := NewToken(claims)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.OK(&app.Token{Token: "Bearer " + token})
}

// Signup runs the signup action.
func (c *AuthController) Signup(ctx *app.SignupAuthContext) error {
	// Payloadバリデーション
	p := ctx.Payload
	if p.DeviceToken != "" && p.InstanceID == "" || p.DeviceToken == "" && p.InstanceID != "" {
		return ctx.BadRequest(goa.ErrBadRequest("DeviceTokenとInstanceIDを片方だけ設定することはできません。"))
	}

	// ユーザーID、メールアドレスが使用可能かを検査
	if c.db.Users.Exists(Keys{"id": p.ID}) {
		return ctx.BadRequest(goa.ErrBadRequest("User ID '" + p.ID + "' is already in use."))
	}
	if c.db.Users.Exists(Keys{"mail.email": p.Mail, "mail.state": STATE_OK}) {
		return ctx.BadRequest(goa.ErrBadRequest("Email address '" + p.Mail + "' is already in use."))
	}

	// パスワードを暗号化
	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// ユーザーをDBに保存
	user := &UserModel{
		ID:       p.ID,
		Password: string(pass),
		Name:     p.Name,
		Icon:     p.Icon,
		Mail: &Mail{
			Email: p.Mail,
			State: STATE_OK,
		},
		Tel: p.Tel,
	}
	// スマホから同時に通知登録されて来た(デバイストークンとインスタンスIDが設定されて来た)とき
	if p.DeviceToken != "" && p.InstanceID != "" {
		// デバイストークンをキーに、インスタンスIDをセット
		user.NotifyTargets = map[string]string{p.DeviceToken: p.InstanceID}
	}
	err = c.db.Users.Save(user, Keys{"mail.email": p.Mail, "mail.state": STATE_PENDING})
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// 登録完了メールを送信
	err = mailer.SendSignUpCompleteMail(user.Mail.Email, user.Name)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// JWTを生成して返す
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
	}
	token, err := NewToken(claims)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}
	return ctx.OK(&app.Token{Token: "Bearer " + token})
}

// VerifyToken runs the verify_token action.
func (c *AuthController) VerifyToken(ctx *app.VerifyTokenAuthContext) error {
	// トークンが使用可能か検査
	u, err := c.db.Users.FindOne(Keys{"mail.token": ctx.Token})
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	return ctx.OK(parser.ToEmailMedia(u.Mail.Email))
}
