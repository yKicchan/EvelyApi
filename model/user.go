package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDB struct {
	*mgo.Database
}

type UserModel struct {
	ID   string  `bson:id`
	Name string  `bson:name`
	Mail string  `bson:mail`
	Tel  *string `bson:tel`
}

// ユーザーの全情報を抽出するセレクタ
var USER_FULL_SELECTOR = bson.M{
	"_id":      0,
	"password": 0,
}

// ユーザーの一部情報のみ抽出するセレクタ
var USER_TINY_SELECTOR = bson.M{
	"_id":      0,
	"id":       1,
	"password": 0,
	"name":     1,
	"mail":     1,
	"tel":      0,
}

/**
 * Usesコレクションを操作するDBオブジェクトを生成して返す
 * @param  db DBオブジェクト
 * @return    Usersコレクション操作用のDBオブジェクト
 */
func NewUserDB(db *mgo.Database) *UserDB {
	return &UserDB{db}
}

/**
 * ログイン認証し、ユーザーが存在したらユーザー情報を返す
 * @param  id       ユーザーID
 * @param  password パスワード
 * @return user     ユーザー情報
 * @return err      検索時に発生したエラー
 */
func (db *UserDB) Authentication(id string, password string) (user *UserModel, err error) {
	query := bson.M{"id": id, "password": password}
	err = db.C("users").Find(query).Select(USER_FULL_SELECTOR).One(&user)
	return
}

/**
 * ユーザーの全情報を返す
 * @param  id   ユーザーID
 * @return user 検索にヒットしたユーザーの情報
 * @return err  検索時に発生したエラー
 */
func (db *UserDB) GetUser(id string) (user *UserModel, err error) {
	query := bson.M{"id": id}
	err = db.C("users").Find(query).Select(USER_FULL_SELECTOR).One(&user)
	return
}
