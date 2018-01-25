package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controllers/parser"
	. "EvelyApi/middleware"
	"EvelyApi/models"
	. "EvelyApi/models/collections"
	. "EvelyApi/models/documents"
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

// UpdateToken runs the update_token action.
func (c *UsersController) Update(ctx *app.UpdateUsersContext) error {
	id, err := GetLoginID(ctx)
	p := ctx.Payload
	if err == nil { // 認証されて来たとき
		// ユーザーの通知トークン情報を更新
		keys := Keys{"id": id}
		u, _ := c.db.Users.FindOne(keys)
		if u.NotifyTargets != nil {
			u.NotifyTargets[p.DeviceToken] = p.InstanceID
		} else {
			u.NotifyTargets = map[string]string{p.DeviceToken: p.InstanceID}
		}
		err = c.db.Users.Save(u, keys)
		if err != nil {
			return ctx.BadRequest(goa.ErrInternal(err))
		}
	} else { // 認証されてこなかったとき
		// ゲストユーザーとしてトークンを保存
		u := &UserModel{
			Mail: &Mail{
				State: STATE_GUEST,
			},
			NotifyTargets: map[string]string{p.DeviceToken: p.InstanceID},
		}
		keys := Keys{"notify_targets." + p.DeviceToken: p.InstanceID}
		err = c.db.Users.Save(u, keys)
		if err != nil {
			return ctx.BadRequest(goa.ErrInternal(err))
		}
	}
	return ctx.OK([]byte("Success!!"))
}
