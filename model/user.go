package model

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserDB struct {
	*mgo.Database
}

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

/**
 * Usesコレクションを操作するDBオブジェクトを生成して返す
 * @param  db DBオブジェクト
 * @return    Usersコレクション操作用のDBオブジェクト
 */
func NewUserDB(db *mgo.Database) *UserDB {
	return &UserDB{db}
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

/**
 * ユーザーを新しく作成する
 * @param  id  ユーザーID
 * @return     作成時に発生したエラー
 */
func (db *UserDB) NewUser(id string) error {
	_, err := db.GetUser(id)
	if err == nil {
		return fmt.Errorf("\"%s\" has already been taken", id)
	}
	user := &UserModel{ID: id}
	return db.C("users").Insert(user)
}

/**
 * ユーザー情報を更新する
 * @param  user ユーザー情報
 * @return      更新時に発生したエラー
 */
func (db *UserDB) SaveUser(user *UserModel) error {
	selector := bson.M{"id": user.ID}
	update := bson.M{"$set": user}
	return db.C("users").Update(selector, update)
}
