package api

import (
	"EvelyApi/app"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
    . "EvelyApi/middleware"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo/bson"
)

// PinsController implements the pins resource.
type PinsController struct {
	*goa.Controller
	db *model.EvelyDB
}

// NewPinsController creates a pins controller.
func NewPinsController(service *goa.Service, db *model.EvelyDB) *PinsController {
	return &PinsController{
		Controller: service.NewController("PinsController"),
		db:         db,
	}
}

// Off runs the off action.
func (c *PinsController) Off(ctx *app.OffPinsContext) error {

	// ユーザー情報から現在のピンの配列を取得
	uid, err := GetLoginID(ctx)
    if err != nil {
        return ctx.BadRequest(goa.ErrBadRequest(err))
    }
	keys := Keys{"id": uid}
	u, _ := c.db.Users.FindOne(keys)

	// ピンするIDを現在のピン配列と比較、あれば削除し、保存
	for _, id := range ctx.Payload.Ids {
		n := indexOf(u.Pins, bson.ObjectIdHex(id))
		if n != -1 {
			u.Pins = append(u.Pins[:n], u.Pins[n+1:]...)
		}
	}
	c.db.Users.Save(u, keys)
	return ctx.OK([]byte("Success!!"))
}

// On runs the on action.
func (c *PinsController) On(ctx *app.OnPinsContext) error {

	// ユーザー情報から現在のピンの配列を取得
	uid, err := GetLoginID(ctx)
    if err != nil {
        return ctx.BadRequest(goa.ErrBadRequest(err))
    }
	keys := Keys{"id": uid}
	u, _ := c.db.Users.FindOne(keys)

	// ピンするIDを現在のピン配列と比較、なければ追加し、保存
	for _, id := range ctx.Payload.Ids {
		if indexOf(u.Pins, bson.ObjectIdHex(id)) == -1 {
			u.Pins = append(u.Pins, bson.ObjectIdHex(id))
		}
	}
	c.db.Users.Save(u, keys)
	return ctx.OK([]byte("Success!!"))
}

/**
 * slice(array)の中の指定の値が存在するインデックスを調べる
 * @param  slice 探す対象のスライス
 * @param  val   探す値
 * @return int   インデックス ないとき-1
 */
func indexOf(slice []bson.ObjectId, val bson.ObjectId) int {
	for i := range slice {
		if slice[i] == val {
			return i
		}
	}
	return -1
}
