package collection

import (
	. "EvelyApi/model/document"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/**
 * PendingUsersコレクションを操作するためのオブジェクト
 * EvelyCollectionを実装している
 */
type PendingUsersCollection struct {
	*mgo.Collection
}

func NewPendingUsersCollection(c *mgo.Collection) *PendingUsersCollection {
	return &PendingUsersCollection{c}
}

/**
 * 認証待ちユーザーを保存(更新)する
 * @param  model 保存(更新)する認証待ちユーザーの情報
 * @param  keys  認証待ちユーザーを特定するキー
 * @return error エラー
 */
func (this *PendingUsersCollection) Save(model EvelyModel, keys Keys) error {
	pu := model.GetPendingUser()
	update := bson.M{"$set": pu}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * 認証待ちユーザーを検索する
 * @param  keys  認証待ちユーザーを特定するキー
 * @param  model 認証待ちユーザーの情報
 * @return err   エラー
 */
func (this *PendingUsersCollection) FindDoc(keys Keys) (EvelyModel, error) {
	pu := &PendingUserModel{}
	err := this.Find(keys.ToQuery()).One(&pu)
	return PendingUser(pu), err
}

/**
 * 認証待ちユーザーを削除する
 * @param  keys  認証待ちユーザーを特定するキー
 * @return error エラー
 */
func (this *PendingUsersCollection) Delete(keys Keys) error {
	return this.Remove(keys.ToQuery())
}
