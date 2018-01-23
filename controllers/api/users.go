package api

import (
	"EvelyApi/app"
	"EvelyApi/controllers/parser"
	"EvelyApi/models"
	. "EvelyApi/models/collections"
	"github.com/goadesign/goa"
)

// UsersController implements the users resource.
type UsersController struct {
	*goa.Controller
	db *models.EvelyDB
}

// NewUsersController creates a users controller.
func NewUsersController(service *goa.Service, db *models.EvelyDB) *UsersController {
	return &UsersController{
		Controller: service.NewController("UsersController"),
		db:         db,
	}
}

// Show runs the show action.
func (c *UsersController) Show(ctx *app.ShowUsersContext) error {
	user, err := c.db.Users.FindOne(Keys{"id": ctx.UserID})
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	return ctx.OK(parser.ToUserMedia(user))
}
