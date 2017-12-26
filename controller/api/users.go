package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/parser"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	"github.com/goadesign/goa"
	"log"
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
	m, err := c.db.Users().FindDoc(Keys{"id": ctx.UserID})
	if err != nil {
		log.Printf("[EvelyApi] %s", err)
		return ctx.NotFound()
	}
    user := m.Make().User
	return ctx.OK(parser.ToUserMedia(user))
}
