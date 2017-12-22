package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/parser"
	"EvelyApi/model"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo"
	"log"
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
	user, err := c.db.GetUser(ctx.UserID)
	if err != nil {
		log.Printf("[EvelyApi] %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(parser.ToUserMedia(user))
	// UsersController_Show: end_implement
}
