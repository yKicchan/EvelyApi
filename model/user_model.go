package model

import (
	"labix.org/v2/mgo/bson"
)

// ユーザーのDBモデル
type UserModel struct {
	ID       string `bson:"id"`
	Password string `bson:"password"`
	Name     string `bson:"name"`
	Mail     Mail   `bson:"mail"`
	Tel      string `bson:"tel"`
}

type Mail struct {
	Email string `bson:"email"`
	State string `bson:"state"`
}

// ユーザーの一部情報のみ抽出するセレクタ
var USER_TINY_SELECTOR = bson.M{
	"_id":  0,
	"id":   1,
	"name": 1,
	"mail": 1,
}
