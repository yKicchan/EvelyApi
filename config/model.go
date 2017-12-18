package config

import (
	"labix.org/v2/mgo/bson"
)

const (
    // コレクション名
	EVENT_COLLECTION        = "Events"
	USER_COLLECTION         = "Users"
	PENDING_USER_COLLECTION = "PendingUsers"

    // 新規登録時トークンの状態
	STATE_AVAILABLE   = "Available"
	STATE_UNAVAILABLE = "UnAvailable"

    // メールアドレスの状態
    STATE_PENDING = "Pending"
    STATE_OK = "OK"
    STATE_BAN = "BAN"
)

// 全情報を取得するセレクタ
var FULL_SELECTOR = bson.M{"_id": 0}
