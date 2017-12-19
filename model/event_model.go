package model

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// イベントのDBモデル
type EventModel struct {
	ID            string       `bson:id`
	Title         string       `bson:title`
	Host          Host         `bson:host`
	Body          string       `bson:body`
	Place         Location     `bson:place`
	Update_Date   time.Time    `bson:update_date`
	Upcoming_Date UpcomingDate `bson:upcoming_date`
	URL           string       `bson:url`
	Mail          string       `bson:mail`
	Tel           string       `bson:tel`
}

// イベント主催者のDBモデル
type Host struct {
	ID   string `bson:id`
	Name string `bson:name`
}

// イベントの開催場所のDBモデル
type Location struct {
	Name    string     `bson:name`
	Lng_Lat [2]float64 `bson:lng_lat`
}

// イベントの開催予定期間のDBモデル
type UpcomingDate struct {
	Start_Date time.Time `bson:start_date`
	End_Date   time.Time `bson:end_date`
}

// イベントの一部情報のみを取得するセレクタ
var EVENT_TINY_SELECTOR = bson.M{
	"_id":           0,
	"id":            1,
	"title":         1,
	"host":          1,
	"place":         1,
	"upcoming_date": 1,
}
