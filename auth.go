package main

import (
	"EvelyApi/app"
	"github.com/goadesign/goa"
)

// AuthController implements the auth resource.
type AuthController struct {
	*goa.Controller
}

// NewAuthController creates a auth controller.
func NewAuthController(service *goa.Service) *AuthController {
	return &AuthController{Controller: service.NewController("AuthController")}
}

// SendMail runs the send_mail action.
func (c *AuthController) SendMail(ctx *app.SendMailAuthContext) error {
	// AuthController_SendMail: start_implement

	// Put your logic here

	return nil
	// AuthController_SendMail: end_implement
}

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	// AuthController_Signin: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthController_Signin: end_implement
}

// Signup runs the signup action.
func (c *AuthController) Signup(ctx *app.SignupAuthContext) error {
	// AuthController_Signup: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthController_Signup: end_implement
}

// VerifyToken runs the verify_token action.
func (c *AuthController) VerifyToken(ctx *app.VerifyTokenAuthContext) error {
	// AuthController_VerifyToken: start_implement

	// Put your logic here

	res := &app.Email{}
	return ctx.OK(res)
	// AuthController_VerifyToken: end_implement
}
