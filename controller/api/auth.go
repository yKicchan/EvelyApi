package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controller/mailer"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/document"
	"context"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"time"
)

/**
 * メールアドレスをレスポンス形式に変換する
 * @param  email メールアドレス
 * @return Email レスポンス形式に変換したメールアドレス
 */
func ToEmailMedia(email string) *app.Email {
	return &app.Email{
		Email: email,
	}
}

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db *model.EvelyDB
}

/**
 * 受け取った任意の情報を付加してトークンを生成する
 * @param  claims トークンに埋め込む情報
 * @return string 生成したトークン
 */
func newToken(claims jwtgo.MapClaims) string {
	// jwt生成
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	token.Claims = claims
	// 秘密鍵読み込み
	pem, err := ioutil.ReadFile("./keys/id_rsa")
	if err != nil {
		log.Fatalf("[EvelyApi] faild to load file: %s", err)
	}
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		log.Fatalf("[EvelyApi] faild to parse private key: %s", err)
	}
	// jwtを暗号化
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("[EvelyApi] failed to sign token: %s", err)
	}
	return signedToken
}

/**
 * JWTを複合して設定したJSON情報を返す
 * @param  ctx    コンテキスト
 * @return claims jwtに埋め込んだjson
 */
func GetJWTClaims(ctx context.Context) (claims jwtgo.MapClaims) {
	token := jwt.ContextJWT(ctx)
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		log.Fatalf("unsupported claims shape")
	}
	return claims
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *model.EvelyDB) *AuthController {
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         db,
	}
}

// SendMail runs the send_mail action.
func (c *AuthController) SendMail(ctx *app.SendMailAuthContext) error {

	// メールアドレスが使用可能か検査
	email := ctx.Payload.Email
	res := c.db.Users().VerifyEmail(email)
	if !res {
		return ctx.BadRequest()
	}

	// トークンを発行する
	claims := jwtgo.MapClaims{
		"scopes":     "api:access",
		"email":      email,
		"created_at": time.Now(),
	}
	token := newToken(claims)

	// メール送信
	subject := "仮登録完了"
	url := "http://localhost:8888/verify_email?token=" + token
	body := "登録用URL: " + url
	err := mailer.SendMail(email, subject, body)
	if !res {
		log.Printf("[EvelyApi] faild to send email: %s", err)
		return ctx.BadRequest()
	}

	// 認証待ちユーザーをDBに保存
	pu := &PendingUserModel{
		Email:     email,
		Token:     token,
		CreatedAt: claims["created_at"].(time.Time),
	}
	keys := Keys{"email": pu.Email}
	err = c.db.PendingUsers().Save(PendingUser(pu), keys)
	if err != nil {
		log.Printf("[EvelyApi] faild to create pending user: %s", err)
		return ctx.BadRequest()
	}
	return ctx.OK([]byte("Success!!"))
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {

	// ユーザーが存在するか検索
	p := ctx.Payload
	m, err := c.db.Users().FindDoc(Keys{"id": p.ID})
	user := m.Make().User
	if err != nil {
		return ctx.BadRequest()
	}

	// IDとパスワードが一致するかを検査
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		log.Printf("[EvelyApi] %s", err)
		return ctx.BadRequest()
	}

	// JWTを生成して返す
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
		"name":   user.Name,
	}
	token := newToken(claims)
	return ctx.OK(&app.Token{Token: "Bearer " + token})
}

// Signup runs the signup action.
func (c *AuthController) Signup(ctx *app.SignupAuthContext) error {

	// ユーザーIDが使用可能かを検査
	p := ctx.Payload
	uc := c.db.Users()
	if !uc.VerifyID(p.ID) {
		log.Printf("[EvelyApi] faild to create user: \"" + p.ID + "\" is already in use.")
		return ctx.BadRequest()
	}

	// パスワードを暗号化
	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[EvelyApi] faild to generate hash: %s", err)
		return ctx.BadRequest()
	}

	// ユーザーをDBに保存
	user := &UserModel{
		ID:       p.ID,
		Password: string(pass),
		Name:     p.Name,
		Mail: &Mail{
			Email: p.Mail,
			State: STATE_OK,
		},
		Tel: p.Tel,
	}
	keys := Keys{"id": user.ID}
	err = uc.Save(User(user), keys)
	if err != nil {
		log.Printf("[EvelyApi] faild to save user: %s", err)
		return ctx.BadRequest()
	}

	// 一時ユーザーを削除する
	err = c.db.PendingUsers().Delete(Keys{"email": user.Mail.Email})
	if err != nil {
		log.Printf("[EvelyApi] faild to delete pending user: %s", err)
	}

	// JWTを生成して返す
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
		"name":   user.Name,
	}
	token := newToken(claims)
	return ctx.OK(&app.Token{Token: "Bearer " + token})
}

// VerifyToken runs the verify_token action.
func (c *AuthController) VerifyToken(ctx *app.VerifyTokenAuthContext) error {
	// トークンが使用可能か検査
	model, err := c.db.PendingUsers().FindDoc(Keys{"token": ctx.Token})
	pu := model.Make().PendingUser
	if err != nil {
		log.Printf("[EvelyApi] faild to verify email: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(ToEmailMedia(pu.Email))
}
