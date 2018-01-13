package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/parser"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *model.EvelyDB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *model.EvelyDB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	user, err := c.db.Users.FindOne(Keys{"id": ctx.UserID})
	if err != nil {
		return ctx.NotFound(err)
	}
	return ctx.OK(parser.ToUserMedia(user))
}
