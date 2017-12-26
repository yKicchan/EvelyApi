package api

import (
	"EvelyApi/app"
	"EvelyApi/controller/parser"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/document"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo/bson"
	"log"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *model.EvelyDB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *model.EvelyDB) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *EventsController) Create(ctx *app.CreateEventsContext) error {

	// JWTからユーザー情報を取得
	claims := GetJWTClaims(ctx)
	user := &UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}

	// イベントを作成
	p := ctx.Payload
	event := parser.ToEventModel(p, bson.NewObjectId(), user)
    keys := Keys{"_id": event.ID}
	err := c.db.Events().Save(Event(event), keys)
	if err != nil {
		log.Printf("[EvelyApi] faild to save event: %s", err)
		return ctx.BadRequest()
	}

	return ctx.Created(parser.ToEventMedia(event))
}

// Delete runs the delete action.
func (c *EventsController) Delete(ctx *app.DeleteEventsContext) error {

	// JWTからユーザーIDを取得、削除権限があるかを検査
	claims := GetJWTClaims(ctx)
	userID := claims["id"].(string)
	if userID != ctx.UserID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden()
	}

	// イベントを削除する
	err := c.db.Events().Delete(Keys{"_id": ctx.EventID})
	if err != nil {
		log.Printf("[EvelyApi] failed to delete event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK([]byte("Seccess!!"))
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {

	// 条件と一致するイベントを複数検索
	events, err := c.db.Events().FindEvents(
		ctx.Limit,
		ctx.Offset,
		WithKeyword(ctx.Keyword),
		WithUserID(ctx.UserID),
	)
	if err != nil {
		log.Printf("[EvelyApi] faild to search events: %s", err)
		return ctx.NotFound()
	}

	// イベント情報をレスポンス形式に変換して返す
	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventTinyMedia(events[i])
	}
	return ctx.OKTiny(res)
}

// Show runs the show action.
func (c *EventsController) Show(ctx *app.ShowEventsContext) error {
	// IDと一致するイベントを検索
	model, err := c.db.Events().FindDoc(Keys{"_id": ctx.EventID})
    event := model.Make().Event
	if err != nil {
		log.Printf("[EvelyApi] faild to find event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(parser.ToEventMedia(event))
}

// Modify runs the modify action.
func (c *EventsController) Modify(ctx *app.ModifyEventsContext) error {
	// JWTからユーザー情報を取得する
	claims := GetJWTClaims(ctx)
	user := &UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}
	// 編集権限があるかを判定
	if user.ID != ctx.UserID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden()
	}

	// DBのイベント情報を更新
	event := parser.ToEventModel(ctx.Payload, bson.ObjectIdHex(ctx.EventID), user)
    keys := Keys{"_id": event.ID}
	err := c.db.Events().Save(Event(event), keys)
	if err != nil {
		log.Printf("[EvelyApi] faild to save event: %s", err)
		return ctx.NotFound()
	}
	return ctx.OK(parser.ToEventMedia(event))
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	return ctx.OK([]byte("現在実装中！\n完成までしばし待たれよ。"))
}
