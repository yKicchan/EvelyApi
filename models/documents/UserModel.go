package documents

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// ユーザーの全情報を抽出するセレクタ
var USER_DEFAULT_SELECTOR = bson.M{
	"_id":       0,
	"id":        1,
	"password":  1,
	"name":      1,
	"icon":      1,
	"mail":      1,
	"tel":       1,
	"pins":      1,
	"create_at": 1,
}

// ユーザーの一部情報のみ抽出するセレクタ
var USER_TINY_SELECTOR = bson.M{
	"_id":  0,
	"id":   1,
	"name": 1,
	"icon": 1,
}

// ユーザーのDBモデル
type UserModel struct {
	ID            string            `bson:"id"`
	Password      string            `bson:"password"`
	Name          string            `bson:"name"`
	Icon          string            `bson:"icon"`
	Mail          *Mail             `bson:"mail"`
	Tel           string            `bson:"tel"`
	Pins          []bson.ObjectId   `bson:"pins"`
	NotifyTargets map[string]string `bson:"notify_targets"`
	CreatedAt     time.Time         `bson:"created_at"`
}

type Mail struct {
	Email string `bson:"email"`
	State string `bson:"state"`
	Token string `bson:"token"`
}
