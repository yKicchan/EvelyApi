package controller

import (
	"EvelyApi/app"
	"EvelyApi/model"
	"github.com/goadesign/goa"
	mgo "gopkg.in/mgo.v2"
	"log"
)

/**
 * ユーザー情報をAPIのレスポンス形式に変換する
 * @param  u ユーザー情報
 * @return   レスポンス形式に変換したユーザー情報
 */
func ToUserMedia(u *model.UserModel) *app.User {
	return &app.User{
		ID:   u.ID,
		Name: u.Name,
		Mail: u.Mail,
		Tel:  u.Tel,
	}
}

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
	return ctx.OK(ToUserMedia(user))
	// UsersController_Show: end_implement
}
