package collections

import (
	. "EvelyApi/models/documents"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/**
 * Usersコレクションを操作するためのオブジェクト
 */
type UsersCollection struct {
	*mgo.Collection
}

func NewUsersCollection(c *mgo.Collection) *UsersCollection {
	return &UsersCollection{c}
}

/**
 * ユーザーを保存(更新)する
 * @param  user  保存(更新)するユーザーの情報
 * @param  keys  ユーザーを特定するキー
 * @return error エラー
 */
func (this *UsersCollection) Save(user *UserModel, keys Keys) error {
	update := bson.M{"$set": user}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * ユーザーを検索する
 * @param  keys ユーザーを特定するキー
 * @return user ユーザーの情報
 * @return err  エラー
 */
func (this *UsersCollection) FindOne(keys Keys) (u *UserModel, err error) {
	err = this.Find(keys.ToQuery()).One(&u)
	return
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
 * ユーザーが存在しているかを判定する
 * @param  keys ユーザーを特定するキー
 * @return bool true: 存在している, false: 存在していない
 */
func (this *UsersCollection) Exists(keys Keys) bool {
	_, err := this.FindOne(keys)
	return err == nil
}

// /**
//  * ユーザーをオプションから検索する
//  * @param  opt     検索オプション
//  * @return users   検索結果のユーザー配列
//  * @return err     検索時に発生したエラー
//  */
// func (this *UsersCollection) FindUsers(opt FindOptions) (users []*UserModel, err error) {
// 	// 検索オプションの内容からクエリを作成
// 	var query bson.M
//
// 	// 検索条件からイベントを検索
// 	q := this.Find(query).Select(USER_TINY_SELECTOR)
//
//     // 除外件数を指定
// 	if opt.IsOffsetSet() {
// 		q = q.Skip(opt.GetOffset())
// 	}
// 	// 検索件数の上限を指定
// 	if opt.IsLimitSet() {
// 		q = q.Limit(opt.GetLimit())
// 	}
// 	// 結果を返す
// 	err = q.All(&users)
// 	return
// }
