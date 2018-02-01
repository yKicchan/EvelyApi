package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controllers/parser"
	. "EvelyApi/middleware"
	"EvelyApi/models"
	. "EvelyApi/models/collections"
	. "EvelyApi/models/documents"
	jwtgo "github.com/dgrijalva/jwt-go"
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

// Modify runs the modify action.
func (c *UsersController) Modify(ctx *app.ModifyUsersContext) error {
	id, err := GetLoginID(ctx)
	if err != nil {
		return ctx.Unauthorized(goa.ErrUnauthorized(err))
	}
	keys := Keys{"id": id}
	u, _ := c.db.Users.FindOne(keys)
	p := ctx.Payload
	if p.Name != "" {
		u.Name = p.Name
	}
	if p.Icon != "" {
		u.Icon = p.Icon
	}
	if p.Email != "" {
		u.Mail.Email = p.Email
	}
	if p.Tel != "" {
		u.Tel = p.Tel
	}
	err = c.db.Users.Save(u, keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}
	// JWTを生成して返す
	claims := jwtgo.MapClaims{
		"scopes": "api:access",
		"id":     id,
	}
	token, err := NewToken(claims)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}
	return ctx.OK(&app.Token{Token: "Bearer " + token})
}

// ModifyToken runs the modify_token action.
func (c *UsersController) ModifyToken(ctx *app.ModifyTokenUsersContext) error {
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

// Setting runs the setting action.
func (c *UsersController) Setting(ctx *app.SettingUsersContext) error {
	// 設定変更
	id, err := GetLoginID(ctx)
	if err != nil {
		return ctx.Unauthorized(goa.ErrUnauthorized(err))
	}
	keys := Keys{"id": id}
	u, _ := c.db.Users.FindOne(keys)
	p := ctx.Payload
	// カテゴリ
	if len(p.Preferences) > 0 {
		u.Preferences = make([]string, len(p.Preferences))
		copy(u.Preferences, p.Preferences)
	}
	err = c.db.Users.Save(u, keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}
	return ctx.OK([]byte("Success!!"))
}
