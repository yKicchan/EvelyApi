package document

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// イベントの一部情報のみを取得するセレクタ
var EVENT_TINY_SELECTOR = bson.M{
	"_id":          1,
	"id":           1,
	"title":        1,
	"host":         1,
	"plans":        1,
	"notice_range": 1,
	"scope":        1,
	"open_flg":     1,
	"update_date":  1,
}

// イベントのDBモデル
type EventModel struct {
	ID          bson.ObjectId    `bson:"_id"`
	Title       string    `bson:"title"`
	Body        string    `bson:"body"`
	Host        *Host     `bson:"host"`
	Mail        string    `bson:"mail"`
	Tel         string    `bson:"tel"`
	URL         string    `bson:"url"`
	Plans       []*Plan   `bson:"plans"`
	NoticeRange int       `bson:"notice_range"`
	Scope       string    `bson:"scope"`
	OpenFlg     bool      `bson:"open_flg"`
	UpdateDate  time.Time `bson:"update_date"`
	CreatedAt   time.Time `bson:"created_at"`
}

// イベント主催者のDBモデル
type Host struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
}

// イベントの開催予定情報のDBモデル
type Plan struct {
	Location     *Location     `bson:"location"`
	UpcomingDate *UpcomingDate `bson:"upcoming_date"`
}

// イベントの開催場所のDBモデル
type Location struct {
	Name   string     `bson:"name"`
	LngLat [2]float64 `bson:"lng_lat"`
}

// イベントの開催予定期間のDBモデル
type UpcomingDate struct {
	StartDate time.Time `bson:"start_date"`
	EndDate   time.Time `bson:"end_date"`
}
