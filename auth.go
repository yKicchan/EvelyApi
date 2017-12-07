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

// Signin runs the signin action.
func (c *AuthController) Signin(ctx *app.SigninAuthContext) error {
	// AuthController_Signin: start_implement

	// Put your logic here

	res := &app.Token{}
	return ctx.OK(res)
	// AuthController_Signin: end_implement
}
