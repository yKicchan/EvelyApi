package controller

import (
	"EvelyApi/app"
	"github.com/goadesign/goa"
)

// ActionsController implements the actions resource.
type ActionsController struct {
	*goa.Controller
}

// NewActionsController creates a actions controller.
func NewActionsController(service *goa.Service) *ActionsController {
	return &ActionsController{Controller: service.NewController("ActionsController")}
}

// Ping runs the ping action.
func (c *ActionsController) Ping(ctx *app.PingActionsContext) error {
	// ActionsController_Ping: start_implement

	// Put your logic here
	res := []byte("pong")

	// ActionsController_Ping: end_implement
	return ctx.OK(res)
}
