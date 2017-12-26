package collection

import (
	. "EvelyApi/model/document"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/**
 * Usersコレクションを操作するためのオブジェクト
 * EvelyCollectionを実装している
 */
type UsersCollection struct {
	*mgo.Collection
}

func NewUsersCollection(c *mgo.Collection) *UsersCollection {
    return &UsersCollection{c}
}

/**
 * ユーザーを保存(更新)する
 * @param  model 保存(更新)するユーザーの情報
 * @param  keys  ユーザーを特定するキー
 * @return error エラー
 */
func (this *UsersCollection) Save(model EvelyModel, keys Keys) error {
    u := model.Make().User
	update := bson.M{"$set": u}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * ユーザーを検索する
 * @param  keys  ユーザーを特定するキー
 * @param  model ユーザーの情報
 * @return err   エラー
 */
func (this *UsersCollection) FindDoc(keys Keys) (EvelyModel, error) {
    u := &UserModel{}
	err := this.Find(keys.ToQuery()).One(&u)
	return User(u), err
}

/**
 * ユーザーを削除する
 * @param  keys  ユーザーを特定するキー
 * @return error エラー
 */
func (this *UsersCollection) Delete(keys Keys) error {
	return this.Remove(keys.ToQuery())
}

/**
 * ユーザーIDが使用可能かを判定する
 * @param  id         ユーザーID
 * @return true|false 使用可能 | 使用不可(すでに使われている)
 */
func (this *UsersCollection) VerifyID(id string) bool {
	_, err := this.FindDoc(Keys{"id": id})
	return err != nil
}

/**
 * メールアドレスが使用可能かを判定する
 * @param  email      メールアドレス
 * @return true|false 使用可能 | 使用不可(すでに使われている)
 */
func (this *UsersCollection) VerifyEmail(email string) bool {
	_, err := this.FindDoc(Keys{"mail.email": email})
	return err != nil
}
