package model

import (
	"labix.org/v2/mgo/bson"
	"log"
)

/**
 * メールアドレスがすでに使われているかを判定する
 * @param  email メールアドレス
 * @return bool  true: 使用可能, false: 使用不可(すでに使われている)
 */
func (db *UserDB) VerifyEmail(email string) bool {
	query := bson.M{"mail": email}
	user := &UserModel{}
	err := db.C("users").Find(query).One(user)
	log.Printf("[EvelyApi] user: %v", user)
	return err != nil
}

/**
 * PendingUserを作成する(すでにEmailが存在する場合は更新)
 * @param  pendingUser 認証待ちユーザー
 * @return error       クエリ実行時のエラー
 */
func (db *UserDB) CreatePendingUser(pendingUser *PendingUserModel) error {
	selector := bson.M{"email": pendingUser.Email}
	info, err := db.C("PendingUsers").Upsert(selector, pendingUser)
	log.Printf("[EvelyApi] upsert info: %v", info)
	return err
}

/**
 * 認証待ちユーザーを削除する
 * @param  email 認証待ちユーザーだったメールアドレス
 * @return error クエリ実行時のエラー
 */
func (db *UserDB) DeletePendingUser(email string) error {
	selector := bson.M{"email": email}
	return db.C("PendingUsers").Remove(selector)
}

/**
 * 認証待ちユーザーのトークン状態を確認する
 * @param  token  トークン
 * @return string トークン状態を示すメッセージ["Available", "UnAvailable"]
 */
func (db *UserDB) GetTokenState(token string) string {
	query := bson.M{"token": token}
	pendingUser := &PendingUserModel{}
	err := db.C("PendingUsers").Find(query).One(&pendingUser)
	return map[bool]string{true: STATE_AVAILABLE, false: STATE_UNAVAILABLE}[err == nil]
}
