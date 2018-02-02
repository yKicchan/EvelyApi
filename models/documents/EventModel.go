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
	"files":       1,
	"categorys":   1,
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
	"image":     1,
	"title":     1,
	"categorys": 1,
	"host":      1,
	"schedules": 1,
	"reviews":   1,
}

// イベントのDBモデル
type EventModel struct {
	ID          bson.ObjectId   `bson:"_id,omitempty"`
	Image       string          `bson:"image,omitempty"`
	Title       string          `bson:"title,omitempty"`
	Body        string          `bson:"body,omitempty"`
	Files       []string        `bson:"files,omitempty"`
	Categorys   []string        `bson:"categorys,omitempty"`
	Host        *Host           `bson:"host,omitempty"`
	Mail        string          `bson:"mail,omitempty"`
	Tel         string          `bson:"tel,omitempty"`
	URL         string          `bson:"url,omitempty"`
	Schedules   []*Schedule     `bson:"schedules,omitempty"`
	NoticeRange int             `bson:"notice_range,omitempty"`
	Scope       string          `bson:"scope,omitempty"`
	OpenFlg     bool            `bson:"open_flg,omitempty"`
	Reviews     []bson.ObjectId `bson:"reviews,omitempty"`
	UpdateDate  time.Time       `bson:"update_date,omitempty"`
	CreatedAt   time.Time       `bson:"created_at,omitempty"`
}

// イベント主催者のDBモデル
type Host struct {
	ID   string `bson:"id,omitempty"`
	Name string `bson:"name,omitempty"`
	Icon string `bson:"icon,omitempty"`
}

// イベントの開催予定情報のDBモデル
type Schedule struct {
	Location     *Location     `bson:"location,omitempty"`
	UpcomingDate *UpcomingDate `bson:"upcoming_date,omitempty"`
}

// イベントの開催場所のDBモデル
type Location struct {
	Name   string     `bson:"name,omitempty"`
	LngLat [2]float64 `bson:"lng_lat,omitempty"`
}

const (
	LNG = iota
	LAT
)

// イベントの開催予定期間のDBモデル
type UpcomingDate struct {
	StartDate time.Time `bson:"start_date,omitempty"`
	EndDate   time.Time `bson:"end_date,omitempty"`
}
