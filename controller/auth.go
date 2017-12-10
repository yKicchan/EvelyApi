package controller

import (
	"EvelyApi/app"
	"EvelyApi/model"
	"context"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware/security/jwt"
	"golang.org/x/crypto/bcrypt"
	mgo "gopkg.in/mgo.v2"
	"io/ioutil"
	"log"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db         *model.UserDB
}

func newToken(user *model.UserModel) *app.Token {
	// jwt生成
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	token.Claims = jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
		"name":   user.Name,
	}
	// 秘密鍵読み込み
	pem := loadPrivateKey()
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		log.Fatalf("[EvelyApi] faild to parse private key: %s", err)
	}
	// jwtを暗号化
	signedToken, err := token.SignedString(privateKey)
	if err != nil {
		log.Fatalf("[EvelyApi] failed to sign token: %s", err)
	}
	return &app.Token{Token: "Bearer " + signedToken}
}

/**
 * 秘密鍵を読み込む
 * @return 秘密鍵
 */
func loadPrivateKey() []byte {
	pem, err := ioutil.ReadFile("./keys/id_rsa")
	if err != nil {
		log.Fatalf("[EvelyApi] faild to load file: %s", err)
	}
	return pem
}

/**
 * JWTを複合してログインしているユーザーの情報を取得する
 * @param  ctx  コンテキスト
 * @return user ログインしているユーザーの情報
 */
func GetLoginUser(ctx context.Context) (user *model.UserModel) {
	token := jwt.ContextJWT(ctx)
	claims, ok := token.Claims.(jwtgo.MapClaims)
	if !ok {
		log.Fatalf("unsupported claims shape")
	}
	user = &model.UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}
	return user
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *mgo.Database) *AuthController {
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         model.NewUserDB(db),
	}
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	// AuthController_Signin: start_implement

	// Put your logic here
	// ログイン認証
	p := ctx.Payload
	user, err := c.db.GetUser(p.ID)
	if err != nil {
		return ctx.Unauthorized()
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(p.Password))
	if err != nil {
		log.Printf("[EvelyApi] %s", err)
		return ctx.Unauthorized()
	}
	return ctx.OK(newToken(user))
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
		Mail:     p.Mail,
		Tel:      p.Tel,
	}
	err = c.db.SaveUser(user)
	if err != nil {
		log.Printf("[EvelyApi] faild to save user: %s", err)
		return ctx.BadRequest()
	}

	return ctx.OK(newToken(user))
	// AuthController_Signup: end_implement
}
