package model

import (
	"labix.org/v2/mgo/bson"
)

// コレクション名を定数化
const (
	EVENT_COLLECTION        = "events"
	USER_COLLECTION         = "users"
	PENDING_USER_COLLECTION = "pending_users"
)

// 全情報を取得するセレクタ
var FULL_SELECTOR = bson.M{"_id": 0}

// トークン状態を表す定数
const (
	STATE_AVAILABLE   = "Available"
	STATE_UNAVAILABLE = "UnAvailable"
)
