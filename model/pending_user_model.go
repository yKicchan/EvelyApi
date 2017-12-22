package model

import (
	"time"
)

// 認証待ちユーザーモデル
type PendingUserModel struct {
	Email     string    `bson:"email"`
	Token     string    `bson:"token"`
	CreatedAt time.Time `bson:"created_at"`
}
