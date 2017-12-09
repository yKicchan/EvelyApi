package controller

import (
	"EvelyApi/app"
	"EvelyApi/model"
	"github.com/goadesign/goa"
	mgo "gopkg.in/mgo.v2"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *model.UserDB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *mgo.Database) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         model.NewUserDB(db),
	}
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	// UsersController_Show: start_implement

	// Put your logic here

	res := &app.User{}
	return ctx.OK(res)
	// UsersController_Show: end_implement
}
