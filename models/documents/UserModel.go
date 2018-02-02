package documents

import (
	"labix.org/v2/mgo/bson"
	"time"
)

// ユーザーの全情報を抽出するセレクタ
var USER_DEFAULT_SELECTOR = bson.M{
	"_id":         0,
	"id":          1,
	"password":    1,
	"name":        1,
	"icon":        1,
	"mail":        1,
	"tel":         1,
	"pins":        1,
	"preferences": 1,
	"create_at":   1,
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
	ID            string            `bson:"id,omitempty"`
	Password      string            `bson:"password,omitempty"`
	Name          string            `bson:"name,omitempty"`
	Icon          string            `bson:"icon,omitempty"`
	Mail          *Mail             `bson:"mail,omitempty"`
	Tel           string            `bson:"tel,omitempty"`
	Pins          []bson.ObjectId   `bson:"pins,omitempty"`
	NotifyTargets map[string]string `bson:"notify_targets,omitempty"`
	Preferences   []string          `bson:"preferences,omitempty"`
	CreatedAt     time.Time         `bson:"created_at,omitempty"`
}

type Mail struct {
	Email string `bson:"email,omitempty"`
	State string `bson:"state,omitempty"`
	Token string `bson:"token,omitempty"`
}
