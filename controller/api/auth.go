package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/mailer"
	"EvelyApi/model"
	"context"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/crypto/bcrypt"
	"labix.org/v2/mgo"
	"io/ioutil"
	"log"
	"time"
)

/**
 * トークンの状態メッセージをレスポンス形式に変換する
 * @param  msg        トークンの状態を示すメッセージ
 * @return TokenState レスポンス形式に変換した情報
 */
func ToTokenStateMedia(msg string) *app.TokenState {
	return &app.TokenState{
		State: msg,
	}
}

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db *model.UserDB
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
func NewAuthController(service *goa.Service, db *mgo.Database) *AuthController {
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         model.NewUserDB(db),
	}
}

// SendMail runs the send_mail action.
func (c *AuthController) SendMail(ctx *app.SendMailAuthContext) error {
	// AuthController_SendMail: start_implement
	// メールアドレスが使用可能か検査
	email := ctx.Payload.Email
	res := c.db.VerifyEmail(email)
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
	// 認証待ちユーザーとしてDBに登録
	pendingUser := &model.PendingUserModel{
		Email:     email,
		Token:     token,
		CreatedAt: claims["created_at"].(time.Time),
	}
	err = c.db.CreatePendingUser(pendingUser)
	if err != nil {
		log.Printf("[EvelyApi] faild to create pending user: %s", err)
		return ctx.BadRequest()
	}
	return ctx.OK([]byte("Success!!"))
	// AuthController_SendMail: end_implement
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	// AuthController_Signin: start_implement

	// Put your logic here
	// ログイン認証
	p := ctx.Payload
	user, err := c.db.GetUser(p.ID)
	if err != nil {
		return ctx.BadRequest()
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		log.Printf("[EvelyApi] %s", err)
		return ctx.BadRequest()
	}
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
		"name":   user.Name,
	}
	token := newToken(claims)
	return ctx.OK(&app.Token{Token: "Bearer " + token})
	// AuthController_Signin: end_implement
}

// Signup runs the signup action.
func (c *AuthController) Signup(ctx *app.SignupAuthContext) error {
	// AuthController_Signup: start_implement

	// Put your logic here
	p := ctx.Payload
	err := c.db.NewUser(p.ID)
	if err != nil {
		log.Printf("[EvelyApi] faild to create user: %s", err)
		return ctx.BadRequest()
	}

	pass, err := bcrypt.GenerateFromPassword([]byte(p.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("[EvelyApi] faild to generate hash: %s", err)
		return ctx.BadRequest()
	}
	log.Print("[Password]%s", string(pass))
	user := &model.UserModel{
		ID:       p.ID,
		Password: string(pass),
		Name:     p.Name,
		Mail: model.Mail{
			Email:  p.Mail,
			Status: 1,
		},
		Tel: p.Tel,
	}
	err = c.db.SaveUser(user)
	if err != nil {
		log.Printf("[EvelyApi] faild to save user: %s", err)
		return ctx.BadRequest()
	}
	_ = c.db.DeletePendingUser(user.Mail.Email)
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
		"name":   user.Name,
	}
	token := newToken(claims)
	return ctx.OK(&app.Token{Token: "Bearer " + token})
	// AuthController_Signup: end_implement
}

// VerifyToken runs the verify_token action.
func (c *AuthController) VerifyToken(ctx *app.VerifyTokenAuthContext) error {
	// AuthController_VerifyToken: start_implement

	// Put your logic here
	msg := c.db.GetTokenState(ctx.Token)

	return ctx.OK(ToTokenStateMedia(msg))
	// AuthController_VerifyToken: end_implement
}
