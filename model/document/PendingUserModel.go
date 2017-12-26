package document

import (
	"labix.org/v2/mgo/bson"
	"time"
)

var PENDING_USER_FULL_SELECTOR = bson.M{"_id": 0}

// 認証待ちユーザーモデル
type PendingUserModel struct {
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
}
