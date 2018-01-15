package collection

import (
	. "EvelyApi/config"
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

/**
 * ユーザーをオプションから検索する
 * @param  options 検索オプション
 * @return users   検索結果のユーザー配列
 * @return err     検索時に発生したエラー
 */
func (this *UsersCollection) FindUsers(options ...FindOptions) (users []*UserModel, err error) {
	// 検索オプションを取得、設定
	opt := findOptions{}
	for _, o := range options {
		o(&opt)
	}
	// 検索オプションの内容からクエリを作成
	query := bson.M{}
	// 位置情報検索
	if opt.r > 0 {
		query = bson.M{
			"lng_lat": bson.M{
				"$nearSphere":  []float64{opt.lng, opt.lat},
				"$maxDistance": (float64(opt.r) * DEGREE_PER_METER),
			},
		}
	}
	// 検索条件からイベントを検索
	q := this.Find(query).Select(USER_TINY_SELECTOR)

	// 除外件数を指定
	if opt.offset > 0 {
		q = q.Skip(opt.offset)
	}
	// 検索件数の上限を指定
	if opt.limit > 0 {
		q = q.Limit(opt.limit)
	}
	// 結果を返す
	err = q.All(&users)
	return
}
