package model

import (
	. "EvelyApi/config"
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
	err := db.C(USER_COLLECTION).Find(query).One(user)
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
	info, err := db.C(PENDING_USER_COLLECTION).Upsert(selector, pendingUser)
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
	return db.C(PENDING_USER_COLLECTION).Remove(selector)
}

/**
 * 認証待ちユーザーのトークンが有効か確認する
 * @param  token  トークン
 * @return string メールアドレス
 * @return error  クエリ実行時のエラー
 */
func (db *UserDB) VerifyToken(token string) (string, error) {
	query := bson.M{"token": token}
	pu := &PendingUserModel{}
	err := db.C(PENDING_USER_COLLECTION).Find(query).One(&pu)
	return pu.Email, err
}
