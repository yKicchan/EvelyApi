package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EventDB struct {
	*mgo.Database
}

type EventModel struct {
	ID          int      `bson:id`
	Title       string   `bson:title`
	Description string   `bson:description`
	Category    string   `bson:category`
	Location    Location `bson:location`
}

type Location [2]float64

func NewEventDB(db *mgo.Database) *EventDB {
	return &EventDB{db}
}

// イベント作成
func (db *EventDB) NewEvent() (event *EventModel, err error) {

	// イベントを作成する
	c := db.C("events")
	event.ID, _ = c.Count()
	event.ID++
	err = c.Insert(event)

	return
}

// イベント複数取得
func (db *EventDB) GetEvents(limit int, offset int, category *string, location *Location) (events []EventModel, err error) {

	// クエリ情報からイベントを複数件取得
	c := db.C("events")
	query := bson.M{}
	selector := bson.M{"_id": 0}

	// カテゴリが指定されていたらクエリを更新
	if category != nil {
		query["category"] = category
	}

	// 座標が指定されていたらクエリを更新
	if location != nil {
		query["location"] = &bson.M{
			"$near": location,
			// 500m
			"$maxDistance": 0.0044938568976209516,
		}
	}

	err = c.Find(query).Select(selector).Skip(offset).Limit(limit).All(&events)
	return
}

// イベント取得
func (db *EventDB) GetEvent(id int) (event *EventModel, err error) {

	// IDと一致するイベントを取得
	c := db.C("events")
	query := bson.M{"id": id}
	selector := bson.M{"_id": 0}
	err = c.Find(query).Select(selector).One(&event)
	return
}

// イベント保存
func (db *EventDB) SaveEvent(event *EventModel) error {

	// イベント更新
	c := db.C("events")
	selector := bson.M{"id": event.ID}
	update := bson.M{"$set": event}
	err := c.Update(selector, update)
	return err
}

// イベント削除
func (db *EventDB) DeleteEvent(id int) error {

	// イベント削除
	c := db.C("events")
	selector := bson.M{"id": id}
	err := c.Remove(selector)
	return err
}
