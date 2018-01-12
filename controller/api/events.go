package api

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controller/parser"
	"EvelyApi/model"
	. "EvelyApi/model/collection"
	. "EvelyApi/model/document"
	"errors"
	"github.com/NaySoftware/go-fcm"
	"github.com/goadesign/goa"
	"labix.org/v2/mgo/bson"
	"log"
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
		return ctx.BadRequest(err)
	}

	return ctx.Created(parser.ToEventMedia(event))
}

// Delete runs the delete action.
func (c *EventsController) Delete(ctx *app.DeleteEventsContext) error {

	// JWTからユーザーIDを取得
	claims := GetJWTClaims(ctx)
	uid := claims["id"].(string)

    // 削除権限があるか(本人のものか)を判定
	eid := bson.ObjectIdHex(ctx.EventID)
	m, err := c.db.Events().FindDoc(Keys{"_id": eid})
    e := m.GetEvent()
    if err != nil {
        return ctx.NotFound(err)
    } else if uid != e.Host.ID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden(errors.New("You do not have permission to delete events."))
	}

	// イベントを削除する
	err = c.db.Events().Delete(Keys{"_id": eid})
	if err != nil {
		log.Printf("[EvelyApi] failed to delete event: %s", err)
		return ctx.BadRequest(err)
	}
	return ctx.OK([]byte("Seccess!!"))
}

// List runs the list action.
func (c *EventsController) List(ctx *app.ListEventsContext) error {

	// 条件と一致するイベントを複数検索
	events, err := c.db.Events().FindEvents(
		WithLimit(ctx.Limit),
		WithOffset(ctx.Offset),
		WithKeyword(ctx.Keyword),
	)
	if err != nil {
		log.Printf("[EvelyApi] faild to search events: %s", err)
		return ctx.BadRequest(err)
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
        m, _ := c.db.Events().FindDoc(Keys{"_id": id})
        events = append(events, m.GetEvent())
    }

    // イベント情報をレスポンス形式に変換して返す
    res := make(app.EventCollection, len(events))
	for i := range events {
		res[i] = parser.ToEventMedia(events[i])
	}
    return ctx.OK(res)
}

// Modify runs the modify action.
func (c *EventsController) Modify(ctx *app.ModifyEventsContext) error {

	// JWTからユーザー情報を取得する
	claims := GetJWTClaims(ctx)
	user := &UserModel{
		ID:   claims["id"].(string),
		Name: claims["name"].(string),
	}
	// 編集権限があるか(本人のものか)を判定
    eid := bson.ObjectIdHex(ctx.EventID)
	m, err := c.db.Events().FindDoc(Keys{"_id": eid})
    e := m.GetEvent()
    if err != nil {
        return ctx.NotFound(err)
    } else if user.ID != e.Host.ID {
		log.Printf("[EvelyApi] permission error")
		return ctx.Forbidden(errors.New("You do not have permission to edit events"))
	}

	// DBのイベント情報を更新
	event := parser.ToEventModel(ctx.Payload, eid, user)
	keys := Keys{"_id": eid}
	err = c.db.Events().Save(Event(event), keys)
	if err != nil {
		log.Printf("[EvelyApi] faild to save event: %s", err)
		return ctx.BadRequest(err)
	}
	return ctx.OK(parser.ToEventMedia(event))
}

// Nearby runs the nearby action.
func (c *EventsController) Nearby(ctx *app.NearbyEventsContext) error {
	return nil
}

// Notify runs the notify action.
func (c *EventsController) Notify(ctx *app.NotifyEventsContext) error {
    p := ctx.Payload
	// 現在地から最大通知範囲より内のイベントを取得
	events, err := c.db.Events().FindEvents(WithLocation(p.Lat, p.Lng, MAX_NOTICE_RANGE))
	if err != nil {
		log.Printf("[EvelyApi] faild to search events: %s", err)
		return ctx.BadRequest(err)
	}

	// ユーザーの位置情報を一時保存
	user := &UserModel{LngLat: [2]float64{p.Lng, p.Lat}}
	keys := Keys{"device_token": p.DeviceToken}
	err = c.db.Users().Save(User(user), keys)
	if err != nil {
		log.Printf("[EvelyApi] failed to save user: %s", err)
		return ctx.BadRequest(err)
	}

	// 近くのイベントの通知範囲内にユーザーが存在するかを調べる
	var nears []*EventModel
	for _, event := range events {
		// 通知範囲(m)を度単位に変換
		r := float64(event.NoticeRange) * DEGREE_PER_METER
		for _, plan := range event.Plans {
			// イベントの通知範囲内にユーザーがいたらそのイベントを保存していく
			users, _ := c.db.Users().FindUsers(WithLocation(plan.Location.LngLat[0], plan.Location.LngLat[1], r), WithDeviceToken(p.DeviceToken))
			if len(users) > 0 {
				nears = append(nears, event)
				break
			}
		}
	}

	// 通知するイベントがなかったときは終了
	if len(nears) == 0 {
		return ctx.OK([]byte("付近にイベントはありませんでした"))
	}

	// 一番近かったイベントを通知内容に設定する
	data := map[string]string{
		"sum": "近くで" + nears[0].Title + "が開催されています！",
	}
	// 他にイベントが複数件あった場合Tipsを設定
	if len(nears) > 1 {
		data["msg"] = "他" + strconv.Itoa(len(nears)) + "件のイベント"
	}

	// デバイストークンを設定しプッシュ通知送信
	ids := []string{
		p.DeviceToken,
	}
	cl := fcm.NewFcmClient(FCM_SERVER_KEY)
	cl.NewFcmRegIdsMsg(ids, data)
	status, err := cl.Send()
	if err != nil {
		log.Printf("[EvelyApi] faild to send FCM :%s", err)
        ctx.BadRequest(err)
	} else {
		status.PrintResults()
	}
	return nil
}

// Update runs the update action.
func (c *EventsController) Update(ctx *app.UpdateEventsContext) error {
	return ctx.OK([]byte("現在実装中"))
}
