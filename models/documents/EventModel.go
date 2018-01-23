package documents

import (
	"labix.org/v2/mgo/bson"
	"time"
)

var EVENT_DEFAULT_SELECTOR = bson.M{
	"_id":         1,
	"image":       1,
	"title":       1,
	"body":        1,
	"files": 1,
	"host":        1,
	"mail":        1,
	"tel":         1,
	"url":         1,
	"schedules":   1,
	"reviews":     1,
	"update_date": 1,
	"created_at":  1,
}

// イベントの一部情報のみを取得するセレクタ
var EVENT_TINY_SELECTOR = bson.M{
	"_id":       1,
	"image": 1,
	"title":     1,
	"host":      1,
	"schedules": 1,
	"reviews":   1,
}

// イベントのDBモデル
type EventModel struct {
	ID          bson.ObjectId   `bson:"_id"`
	Image string `bson:"image"`
	Title       string          `bson:"title"`
	Body        string          `bson:"body"`
	Files []string `bson:"files"`
	Host        *Host           `bson:"host"`
	Mail        string          `bson:"mail"`
	Tel         string          `bson:"tel"`
	URL         string          `bson:"url"`
	Schedules   []*Schedule     `bson:"schedules"`
	NoticeRange int             `bson:"notice_range"`
	Scope       string          `bson:"scope"`
	OpenFlg     bool            `bson:"open_flg"`
	Reviews     []bson.ObjectId `bson:"reviews"`
	UpdateDate  time.Time       `bson:"update_date"`
	CreatedAt   time.Time       `bson:"created_at"`
}

// イベント主催者のDBモデル
type Host struct {
	ID   string `bson:"id"`
	Name string `bson:"name"`
	Icon string `bson:"icon"`
}

// イベントの開催予定情報のDBモデル
type Schedule struct {
	Location     *Location     `bson:"location"`
	UpcomingDate *UpcomingDate `bson:"upcoming_date"`
}

// イベントの開催場所のDBモデル
type Location struct {
	Name   string     `bson:"name"`
	LngLat [2]float64 `bson:"lng_lat"`
}

const (
	LNG = iota
	LAT
)

// イベントの開催予定期間のDBモデル
type UpcomingDate struct {
	StartDate time.Time `bson:"start_date"`
	EndDate   time.Time `bson:"end_date"`
}
