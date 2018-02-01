package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controllers/parser"
	. "EvelyApi/middleware"
	"EvelyApi/models"
	. "EvelyApi/models/collections"
	. "EvelyApi/models/collections/findOptions"
	. "EvelyApi/models/documents"
	"github.com/NaySoftware/go-fcm"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo/bson"
	"math"
	"strconv"
)

// EventsController implements the events resource.
type EventsController struct {
	*goa.Controller
	db *models.EvelyDB
}

// NewEventsController creates a events controller.
func NewEventsController(service *goa.Service, db *models.EvelyDB) *EventsController {
	return &EventsController{
		Controller: service.NewController("EventsController"),
		db:         db,
	}
}

// Create runs the create action.
func (c *EventsController) Create(ctx *app.CreateEventsContext) error {

	// カテゴリの存在チェック
	p := ctx.Payload
	for _, c := range p.Categorys {
		if Categorys.IndexOf(c) == -1 {
			return ctx.BadRequest(goa.ErrBadRequest("Category '" + c + "' dose not exist."))
		}
	}

	// JWTからユーザー情報を取得
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	user, _ := c.db.Users.FindOne(Keys{"id": uid})

	// イベントを作成
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
	opt := NewFindEventsOption()
	opt.SetLimit(ctx.Limit)
	opt.SetOffset(ctx.Offset)
	opt.SetKeyword(ctx.Keyword)
	if ctx.Category != nil {
		opt.SetCategorys([]string{*ctx.Category}, true)
	}
	events, err := c.db.Events.FindEvents(opt)
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

	// カテゴリの存在チェック
	for _, c := range ctx.Payload.Categorys {
		if Categorys.IndexOf(c) == -1 {
			return ctx.BadRequest(goa.ErrBadRequest("Category '" + c + "' dose not exist."))
		}
	}

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
	uid, err := GetLoginID(ctx)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}

	// IDからイベントを取得する
	opt := NewFindEventsOption()
	opt.SetLimit(ctx.Limit)
	opt.SetOffset(ctx.Offset)
	opt.SetHostID(uid)
	events, err := c.db.Events.FindEvents(opt)
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
	opt := NewFindEventsOption()
	opt.SetLimit(ctx.Limit)
	opt.SetLocation(ctx.Lat, ctx.Lng, ctx.Range)
	if ctx.Category != nil {
		opt.SetCategorys([]string{*ctx.Category}, true)
	}
	results, err := c.db.FindEventsByLocation(opt)
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	}
	// イベント情報をレスポンス形式に変換して返す
	res := make(app.NearbyCollection, len(results))
	for i, r := range results {
		res[i] = &app.Nearby{
			Distance: int(r.Distance),
			Event:    parser.ToEventTinyMedia(r.Event),
		}
	}
	return ctx.OK(res)
}

// Notify runs the notify action.
func (c *EventsController) Notify(ctx *app.NotifyEventsContext) error {

	// 位置情報がイベントの通知範囲内か調べ、通知範囲内だったイベントのみを返す
	contain := func(events []*EventModel, lat, lng float64) (res []*EventModel) {
		square := func(x float64) float64 { return x * x }
		for _, e := range events {
			// 通知範囲(m)を度単位に変換
			r := float64(e.NoticeRange) * DEGREE_PER_METER
			for _, schedule := range e.Schedules {
				distance := math.Sqrt(square(lat-schedule.Location.LngLat[LNG]) + square(lng-schedule.Location.LngLat[LAT]))
				if distance > r {
					res = append(res, e)
					break
				}
			}
		}
		return res
	}

	// イベントの検索オプション
	p := ctx.Payload
	opt := NewFindEventsOption()

	// 通知先のインスタンスIDを設定
	var ids []string
	uid, err := GetLoginID(ctx)
	if err == nil { // 認証されて来たとき
		// 登録されている全てのインスタンスIDを設定
		u, _ := c.db.Users.FindOne(Keys{"id": uid})
		for _, id := range u.NotifyTargets {
			ids = append(ids, id)
		}
		// ユーザーが通知を許可しているカテゴリの中からイベント検索
		opt.SetCategorys(u.Preferences, false)
	} else if p.InstanceID != "" { // 認証されてこなかったとき
		// POST送信されてきたインスタンスIDを設定
		ids = []string{p.InstanceID}
	} else { // 認証されず、インスタンスIDもないとき
		return ctx.BadRequest(goa.ErrBadRequest("認証するか、インスタンスIDを設定してください"))
	}

	// 現在地から最大通知範囲より内のイベントを取得
	opt.SetLocation(p.Lat, p.Lng, MAX_NOTICE_RANGE)
	events, err := c.db.Events.FindEvents(opt)
	if err != nil {
		return ctx.BadRequest(goa.ErrInternal(err))
	}

	// 通知範囲内のイベントだけに絞る
	events = contain(events, p.Lat, p.Lng)
	if len(events) == 0 {
		return ctx.OK([]byte("近くに通知するイベントはありませんでした"))
	}

	// 通知メッージを作成
	// 一番近いイベントをタイトルにする
	data := map[string]string{
		"sum": "近くで" + events[0].Title + "が開催されています！",
	}
	// 他にイベントが複数件あった場合Tipsを設定
	if len(events) > 1 {
		data["msg"] = "他" + strconv.Itoa(len(events)) + "件のイベント"
	}

	// プッシュ通知送信
	cl := fcm.NewFcmClient(FCM_SERVER_KEY)
	cl.NewFcmRegIdsMsg(ids, data)
	status, err := cl.Send()
	if err != nil {
		return ctx.BadRequest(goa.ErrBadRequest(err))
	} else {
		status.PrintResults()
	}
	return ctx.OK([]byte(strconv.Itoa(len(events)) + "件のイベントを検知しました"))
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
