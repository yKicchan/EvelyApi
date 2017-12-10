package model

import (
	"time"
	"strconv"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// イベントのリソースに対するDBオブジェクト
type EventDB struct {
	*mgo.Database
}

// イベントのDBモデル
type EventModel struct {
	ID           string       `bson:id`
	Title        string       `bson:title`
	Host         Host         `bson:host`
	Body         string       `bson:body`
	Place        Location     `bson:place`
	UpdateDate   time.Time    `bson:update_date`
	UpcomingDate UpcomingDate `bson:upcoming_date`
	URL          string       `bson:url`
	Mail         string       `bson:mail`
	Tel          string       `bson:tel`
}

// イベント主催者のDBモデル
type Host struct {
	ID   string `bson:id`
	Name string `bson:name`
}

// イベントの開催場所のDBモデル
type Location struct {
	Name   string     `bson:name`
	LngLat [2]float64 `bson:lng_lat`
}

// イベントの開催予定期間のDBモデル
type UpcomingDate struct {
	StartDate time.Time `bson:start_date`
	EndDate   time.Time `bson:end_date`
}

// イベントの検索オプション
type getEventsOptions struct {
	keyword string
	userID  string
}

// イベントの検索オプションを関数化
type GetEventsOption func(*getEventsOptions)

/**
 * ユーザーIDを検索オプションに設定するクロージャを返す
 * @param  userID ユーザーID
 * @return        クロージャ
 */
func WithUserID(userID string) GetEventsOption {
	return func(ops *getEventsOptions) {
		ops.userID = userID
	}
}

/**
 * キーワードを検索オプションに設定するクロージャを返す
 * @param  keyword キーワード
 * @return         クロージャ
 */
func WithKeyword(keyword string) GetEventsOption {
	return func(ops *getEventsOptions)  {
		ops.keyword = keyword
	}
}

var EVENT_FULL_SELECTOR = bson.M{"_id": 0}

var EVENT_TINY_SELECTOR = bson.M{
	"_id":           0,
	"id":            1,
	"title":         1,
	"host":          1,
	"place":         1,
	"upcoming_date": 1,
}

func NewEventDB(db *mgo.Database) *EventDB {
	return &EventDB{db}
}

/**
 * イベントを新しく作成するためのイベントIDを生成し返す
 * @param  userID  作成者のユーザーID
 * @return eventID 生成したイベントのID
 * @return err     作成時に発生したエラー
 */
func (db *EventDB) NewEvent(userID string, upcomingDate time.Time) (eventID string, err error) {
	y, m, d := upcomingDate.Date()
	date := strconv.Itoa(y) + strconv.Itoa(int(m)) + strconv.Itoa(d)
	query := bson.M{
		"$and": []interface{}{
			bson.M{"id": bson.M{"$regex": bson.RegEx{Pattern: date + `-[0-9]+`}}},
			bson.M{"host.id": userID},
		},
	}
	n, _ := db.C("events").Find(query).Count()
	eventID = date + "-" + strconv.Itoa(n + 1)
	event := &EventModel{
		ID: eventID,
		Host: Host{ID: userID},
	}
	err = db.C("events").Insert(event)
	return
}

/**
 * イベント情報をパラメーターから検索し、複数件返す
 * @param  limit   検索件数
 * @param  offset  除外件数
 * @param  options 検索オプション
 * @return events  複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (db *EventDB) GetEvents(limit, offset int, options ...GetEventsOption) (events []EventModel, err error) {
	// 検索オプションを取得
	opt := getEventsOptions{}
    for _, o := range options {
        o(&opt)
    }
	// 検索オプションの内容からクエリを作成
	query := bson.M{}
	if len(opt.keyword) > 0 {
		regex := bson.M{"$regex": bson.RegEx{Pattern: `.*` + opt.keyword + `.*`, Options: "m"}}
		query = bson.M{
			"$or": []interface{}{
				bson.M{"title": regex},
				bson.M{"body": regex},
				bson.M{"place": regex},
				bson.M{"host.name": regex},
			},
		}
	}
	if len(opt.userID) != 0 {
		if len(opt.keyword) > 0 {
			query = bson.M{
				"$and": []interface{}{
					query,
					bson.M{"host.id": opt.userID},
				},
			}
		} else {
			query = bson.M{"host.id": opt.userID}
		}
	}
	err = db.C("events").Find(query).Select(EVENT_TINY_SELECTOR).Skip(offset).Limit(limit).All(&events)
	return
}

/**
 * イベントの詳しい情報を返す
 * @param  userID  主催者のユーザーID
 * @param  eventID イベントID
 * @return event   イベント情報
 * @return err     検索時に発生したエラー
 */
func (db *EventDB) GetEvent(userID, eventID string) (event *EventModel, err error) {
	query := bson.M{
		"$and": []interface{}{
			bson.M{"id":      eventID},
			bson.M{"host.id": userID},
		},
	}
	err = db.C("events").Find(query).Select(EVENT_FULL_SELECTOR).One(&event)
	return
}

/**
 * イベントを保存する
 * @param  event イベント情報
 * @return err   保存時に発生したエラー
 */
func (db *EventDB) SaveEvent(event *EventModel) error {
	selector := bson.M{
		"$and": []interface{}{
			bson.M{"id":      event.ID},
			bson.M{"host.id": event.Host.ID},
		},
	}
	update := bson.M{"$set": event}
	err := db.C("events").Update(selector, update)
	return err
}

// // イベント削除
// func (db *EventDB) DeleteEvent(id int) error {
//
// 	// イベント削除
// 	c := db.C("events")
// 	selector := bson.M{"id": id}
// 	err := c.Remove(selector)
// 	return err
// }
