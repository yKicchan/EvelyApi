package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controller/mail"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/document"
    "EvelyApi/controller/parser"
	"context"
	"errors"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
	"time"
)

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
	if c.db.Users.Exists(Keys{"mail.email": email, "mail.state": STATE_OK}) {
		return ctx.BadRequest(errors.New("\"" + email + "\" is already in use."))
	}

	// トークンを発行する
	claims := jwtgo.MapClaims{
		"scopes":     "api:access",
		"email":      email,
		"created_at": time.Now(),
	}
	token := newToken(claims)

	// メール送信
	url := "http://localhost:8888/verify_email?token=" + token
	err := mail.SendSignUpMail(email, url)
	if err != nil {
		return ctx.BadRequest(err)
	}

	// 認証待ちユーザーをDBに保存
	u := &UserModel{
		Mail: &Mail{
            Email: email,
            Token: token,
            State: STATE_PENDING,
        },
		CreatedAt: claims["created_at"].(time.Time),
	}
	keys := Keys{"mail.email": u.Mail.Email}
	err = c.db.Users.Save(u, keys)
	if err != nil {
		return ctx.BadRequest(err)
	}
	return ctx.OK([]byte("Success!!"))
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {

	// ユーザーが存在するか検索
	p := ctx.Payload
	user, err := c.db.Users.FindOne(Keys{"id": p.ID})
	if err != nil {
		return ctx.BadRequest(errors.New("The ID and password you entered did not match."))
	}

	// IDとパスワードが一致するかを検査
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		return ctx.BadRequest(errors.New("The ID and password you entered did not match."))
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
	if c.db.Users.Exists(Keys{"id": p.ID}) {
		return ctx.BadRequest(errors.New("User ID '" + p.ID + "' is already in use."))
	}

	// パスワードを暗号化
	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		return ctx.BadRequest(err)
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
    if p.DeviceToken != "" {
        user.DeviceToken = p.DeviceToken
    }
    keys := Keys{"device_token": p.DeviceToken}
	err = c.db.Users.Save(user, keys)
	if err != nil {
		return ctx.BadRequest(err)
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
	u, err := c.db.Users.FindOne(Keys{"mail.token": ctx.Token})
	if err != nil {
		return ctx.NotFound(err)
	}
	return ctx.OK(parser.ToEmailMedia(u.Mail.Email))
}
