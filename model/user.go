package model

import (
    . "EvelyApi/config"
	"fmt"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type UserDB struct {
	*mgo.Database
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
	err = db.C("users").Find(query).Select(FULL_SELECTOR).One(&user)
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
