package model

// ユーザーのDBモデル
type UserModel struct {
	ID       string `bson:id`
	Password string `bson:password`
	Name     string `bson:name`
	Mail     string `bson:mail`
	Tel      string `bson:tel`
}

// ユーザーの全情報を抽出するセレクタ
var USER_FULL_SELECTOR = bson.M{
	"_id":      0,
	"id":       1,
	"password": 1,
	"name":     1,
	"mail":     1,
	"tel":      1,
}

// ユーザーの一部情報のみ抽出するセレクタ
var USER_TINY_SELECTOR = bson.M{
	"_id":  0,
	"id":   1,
	"name": 1,
	"mail": 1,
}
