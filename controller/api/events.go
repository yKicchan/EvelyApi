package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controller/parser"
	. "EvelyApi/middleware"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/collection/findOptions"
	. "EvelyApi/model/document"
	"github.com/NaySoftware/go-fcm"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo/bson"
	"math"
	"strconv"
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
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	user, _ := c.db.Users.FindOne(Keys{"id": uid})

	// イベントを作成
	p := ctx.Payload
	event := parser.ToEventModel(p, bson.NewObjectId(), user)
	keys := Keys{"_id": event.ID}
	err = c.db.Events.Save(event, keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	return ctx.Created(parser.ToEventMedia(event))
}

// Delete runs the delete action.
func (c *EventsController) Delete(ctx *app.DeleteEventsContext) error {

	// JWTからユーザー情報を取得
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// 削除権限があるか(本人のものか)を判定
	eid := bson.ObjectIdHex(ctx.EventID)
	keys := Keys{"_id": eid}
	e, err := c.db.Events.FindOne(keys)
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	} else if uid != e.Host.ID {
		errForbidden := goa.NewErrorClass("forbidden", 403)
		return ctx.Forbidden(errForbidden("You do not have permission to delete events."))
	}

	// イベントを削除する
	err = c.db.Events.Delete(keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.OK([]byte("Seccess!!"))
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {

	// 条件と一致するイベントを複数検索
	opt := NewFindEventsOptions()
	opt.SetLimit(ctx.Limit)
	opt.SetOffset(ctx.Offset)
    var events []*EventModel
    var err error
    if ctx.Keyword != "" {
    	opt.SetKeyword(ctx.Keyword)
    	events, err = c.db.Events.FindEventsByKeyword(opt)
    } else {
        events, err = c.db.Events.FindEvents(opt)
    }
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
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
	var events []*EventModel
	for _, id := range ctx.Ids {
		e, err := c.db.Events.FindOne(Keys{"_id": bson.ObjectIdHex(id)})
		if err == nil {
			events = append(events, e)
		}
	}

	// イベント情報をレスポンス形式に変換して返す
	res := make(app.EventCollection, len(events))
	if len(res) == 0 {
		return ctx.OK(res)
	}
	for i := range events {
		res[i] = parser.ToEventMedia(events[i])
	}
	return ctx.OK(res)
}

// Modify runs the modify action.
func (c *EventsController) Modify(ctx *app.ModifyEventsContext) error {

	// JWTからユーザー情報を取得する
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	user, _ := c.db.Users.FindOne(Keys{"id": uid})

	// 編集権限があるか(本人のものか)を判定
	eid := bson.ObjectIdHex(ctx.EventID)
	keys := Keys{"_id": eid}
	e, err := c.db.Events.FindOne(keys)
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	} else if user.ID != e.Host.ID {
		errForbidden := goa.NewErrorClass("forbidden", 403)
		return ctx.Forbidden(errForbidden("You do not have permission to edit events"))
	}

	// DBのイベント情報を更新
	event := parser.ToEventModel(ctx.Payload, eid, user)
	err = c.db.Events.Save(event, keys)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	return ctx.OK(parser.ToEventMedia(event))
}

// MyList runs the my_list action.
func (c *EventsController) MyList(ctx *app.MyListEventsContext) error {

	// JWTからユーザーIDを取得
	claims := GetJWTClaims(ctx)
    uid := claims["id"].(string)

    // IDからイベントを取得する
    events, err := c.db.Events.FindEvents(
        WithHost(uid),
        WithLimit(ctx.Limit),
        WithOffset(ctx.Offset),
    )
    if err != nil {
        ctx.BadRequest(err)
    }

    // イベント情報をレスポンス形式に変換して返す
	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventTinyMedia(events[i])
	}
	return ctx.OKTiny(res)
}

// Nearby runs the nearby action.
func (c *EventsController) Nearby(ctx *app.NearbyEventsContext) error {
	// パラメーターの位置情報から付近のイベントを検索
	opt := NewFindEventsOptions()
	opt.SetLimit(ctx.Limit)
	opt.SetOffset(ctx.Offset)
	opt.SetLocation(ctx.Lat, ctx.Lng, ctx.Range)
	events, err := c.db.Events.FindEventsByLocation(opt)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	// イベント情報をレスポンス形式に変換して返す
	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventTinyMedia(events[i])
	}
	return ctx.OKTiny(res)
}

// NotifyByInstanceID runs the notify_by_instance_id action.
func (c *EventsController) NotifyByInstanceID(ctx *app.NotifyByInstanceIDEventsContext) error {

	// 現在地から最大通知範囲より内のイベントを取得
	p := ctx.Payload
	opt := NewFindEventsOptions()
	opt.SetLocation(p.Lat, p.Lng, MAX_NOTICE_RANGE)
	events, err := c.db.Events.FindEventsByLocation(opt)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// 近くのイベントの通知範囲内にユーザーが存在するかを調べる
	nearbyEvents := contain(events, p.Lat, p.Lng)

	// 通知するイベントがなかったときは終了
	if len(nearbyEvents) == 0 {
		return nil
	}

	// 通知メッセージを作成
	data := createNotifyMessage(nearbyEvents)

	// インスタンスIDを設定しプッシュ通知送信
	ids := []string{p.InstanceID}
	cl := fcm.NewFcmClient(FCM_SERVER_KEY)
	cl.NewFcmRegIdsMsg(ids, data)
	status, err := cl.Send()
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	} else {
		status.PrintResults()
	}
	return nil
}

// NotifyByUserID runs the notify_by_user_id action.
func (c *EventsController) NotifyByUserID(ctx *app.NotifyByUserIDEventsContext) error {

	// 現在地から最大通知範囲より内のイベントを取得
	p := ctx.Payload
	opt := NewFindEventsOptions()
	opt.SetLocation(p.Lat, p.Lng, MAX_NOTICE_RANGE)
	events, err := c.db.Events.FindEventsByLocation(opt)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// 近くのイベントの通知範囲内にユーザーが存在するかを調べる
	nearbyEvents := contain(events, p.Lat, p.Lng)

	// 通知するイベントがなかったときは終了
	if len(nearbyEvents) == 0 {
		return nil
	}

	// 通知メッセージを作成
	data := createNotifyMessage(nearbyEvents)

	// JWTからユーザーIDを取得する
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// インスタンスIDを設定しプッシュ通知送信
	u, err := c.db.Users.FindOne(Keys{"id": uid})
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	cl := fcm.NewFcmClient(FCM_SERVER_KEY)
	cl.NewFcmRegIdsMsg(u.InstanceIds, data)
	status, err := cl.Send()
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	} else {
		status.PrintResults()
	}
	return nil
}

// Pin runs the pin action.
func (c *EventsController) Pin(ctx *app.PinEventsContext) error {

	// ユーザーのピンしているイベント一覧を取得する
	u, err := c.db.Users.FindOne(Keys{"id": ctx.UserID})
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest("User ID '" + ctx.UserID + "' does not exist"))
	}
	var events []*EventModel
	for i := ctx.Offset; i < len(u.Pins) && i < ctx.Limit; i++ {
		e, err := c.db.Events.FindOne(Keys{"_id": u.Pins[i]})
		if err == nil {
			events = append(events, e)
		}
	}

	// イベント情報をレスポンス形式に変換して返す
	res := make(app.EventTinyCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventTinyMedia(events[i])
	}
	return ctx.OKTiny(res)
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	return ctx.OK([]byte("現在実装中"))
}

/**
 * 受け取った位置情報がイベントの通知範囲内にあるかを調べ、通知範囲内だったイベントのみを返す
 * @param  events       検索対象のイベント
 * @param  lat          緯度
 * @param  lng          経度
 * @return nearbyEvents 通知範囲内にあったイベント
 */
func contain(events []*EventModel, lat, lng float64) (nearbyEvents []*EventModel) {
	square := func(x float64) float64 { return x * x }
	for _, e := range events {
		// 通知範囲(m)を度単位に変換
		r := float64(e.NoticeRange) * DEGREE_PER_METER
		for _, schedule := range e.Schedules {
			distance := math.Sqrt(square(lat-schedule.Location.LngLat[LNG]) + square(lng-schedule.Location.LngLat[LAT]))
			if distance > r {
				nearbyEvents = append(nearbyEvents, e)
				break
			}
		}
	}
	return nearbyEvents
}

/**
 * 通知用のメッセージを生成し、返却する
 * @param  events 近くにあったイベント
 * @return msg    生成した通知用メッセージ
 */
func createNotifyMessage(events []*EventModel) (msg map[string]string) {
	// 一番近かったイベントを通知内容に設定する
	msg["sum"] = "近くで" + events[0].Title + "が開催されています！"
	// 他にイベントが複数件あった場合Tipsを設定
	if len(events) > 1 {
		msg["msg"] = "他" + strconv.Itoa(len(events)) + "件のイベント"
	}
	return msg
}
