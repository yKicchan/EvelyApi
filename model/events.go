package model

import (
	"log"
	"time"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EventDB struct {
	*mgo.Database
}

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

type Host struct {
	ID   string `bson:id`
	Name string `bson:name`
}

type Location struct {
	Name   string     `bson:name`
	LngLat [2]float64 `bson:lng_lat`
}

type UpcomingDate struct {
	StartDate time.Time `bson:start_date`
	EndDate   time.Time `bson:end_date`
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
 * @return eventID 作成するイベントのID
 * @return err     作成時に発生したエラー
 */
func (db *EventDB) NewEvent(userID string) (eventID string, err error) {
	query := bson.M{"host": bson.M{"id": userID}}
	n, _ := db.C("users").Find(query).Count()
	eventID = time.Now().Format("19890225") + "-" + string(n)
	return
}

/**
 * イベント情報をパラメーターから検索し、複数件返す
 * @param  limit   検索件数
 * @param  offset  除外件数
 * @param  keyword 検索キーワード
 * @param  userID  ユーザーID
 * @return events  複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (db *EventDB) GetEvents(limit int, offset int, keyword *string, userID *string) (events []EventModel, err error) {
	query := bson.M{}

	if keyword != nil {
		regex := bson.M{"$regex": bson.RegEx{Pattern: `.*` + *keyword + `.*`, Options: "m"}}
		query = bson.M{
			"$or": []interface{}{
				bson.M{"title": regex},
				bson.M{"body": regex},
				bson.M{"place": regex},
				bson.M{"host": bson.M{"name": regex}},
			},
		}
		log.Printf("[EvelyApi] keyword: %s", *keyword)
	}

	if userID != nil {
		query = bson.M{"host": bson.M{"id": bson.M{"$regex": bson.RegEx{Pattern: *userID, Options: "m"}}}}
		log.Printf("[EvelyApi] userID: %s", *userID)
	}
	log.Print(query)

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
func (db *EventDB) GetEvent(userID string, eventID string) (event *EventModel, err error) {
	query := bson.M{
		"id":   eventID,
		"host": bson.M{"id": userID},
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
		"id":   event.ID,
		"host": bson.M{"id": event.Host.ID},
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
