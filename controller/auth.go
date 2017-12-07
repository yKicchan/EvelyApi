package controller

import (
	"EvelyApi/app"
	"EvelyApi/model"
	"crypto/rsa"
	"fmt"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/goadesign/goa"
	mgo "gopkg.in/mgo.v2"
	"io/ioutil"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
	db         *model.UserDB
	privateKey *rsa.PrivateKey
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service, db *mgo.Database) (*AuthController, error) {
	pem, err := ioutil.ReadFile("./keys/id_rsa")
	if err != nil {
		return nil, err
	}
	privateKey, err := jwtgo.ParseRSAPrivateKeyFromPEM(pem)
	if err != nil {
		return nil, err
	}
	return &AuthController{
		Controller: service.NewController("AuthController"),
		db:         model.NewUserDB(db),
		privateKey: privateKey,
	}, nil
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	// AuthController_Signin: start_implement

	// Put your logic here
	// ログイン認証
	payload := ctx.Payload
	user, err := c.db.Authentication(payload.ID, payload.Password)
	if err != nil {
		return ctx.Unauthorized()
	}

	// jwt生成
	token := jwtgo.New(jwtgo.SigningMethodRS512)
	token.Claims = jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     user.ID,
	}
	signedToken, err := token.SignedString(c.privateKey)
	if err != nil {
		return fmt.Errorf("failed to sign token: %s", err) // internal error
	}

	res := &app.Token{Token: "Bearer " + signedToken}
	return ctx.OK(res)
	// AuthController_Signin: end_implement
}
