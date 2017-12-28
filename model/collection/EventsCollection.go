package collection

import (
	. "EvelyApi/model/document"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

/**
 * Eventsコレクションを操作するためのオブジェクト
 * EvelyCollectionを実装している
 */
type EventsCollection struct {
	*mgo.Collection
}

func NewEventsCollection(c *mgo.Collection) *EventsCollection {
	return &EventsCollection{c}
}

/**
 * イベントを保存(更新)する
 * @param  model 保存(更新)するイベントの情報
 * @param  keys  イベントを特定するキー
 * @return error エラー
 */
func (this *EventsCollection) Save(model EvelyModel, keys Keys) error {
	e := model.GetEvent()
	update := bson.M{"$set": e}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * イベントを検索する
 * @param  keys  イベントを特定するキー
 * @param  model イベントの情報
 * @return err   エラー
 */
func (this *EventsCollection) FindDoc(keys Keys) (EvelyModel, error) {
	e := &EventModel{}
	err := this.Find(keys.ToQuery()).One(&e)
	return Event(e), err
}

/**
 * イベントを削除する
 * @param  keys  イベントを特定するキー
 * @return error エラー
 */
func (this *EventsCollection) Delete(keys Keys) error {
	return this.Remove(keys.ToQuery())
}

// イベントの検索オプション
type findEventsOptions struct {
	keyword string
	userID  string
}

// イベントの検索オプションを関数化
type FindEventsOptions func(*findEventsOptions)

/**
 * ユーザーIDを検索オプションに設定するクロージャを返す
 * @param  id ユーザーID
 * @return    クロージャ
 */
func WithUserID(id string) FindEventsOptions {
	return func(ops *findEventsOptions) {
		ops.userID = id
	}
}

/**
 * キーワードを検索オプションに設定するクロージャを返す
 * @param  keyword キーワード
 * @return         クロージャ
 */
func WithKeyword(keyword string) FindEventsOptions {
	return func(ops *findEventsOptions) {
		ops.keyword = keyword
	}
}

/**
 * イベント情報をパラメーターから検索し、複数件返す
 * @param  limit   検索件数
 * @param  offset  除外件数
 * @param  options 検索オプション
 * @return events  複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEvents(limit, offset int, options ...FindEventsOptions) (events []*EventModel, err error) {
	// 検索オプションを取得
	opt := findEventsOptions{}
	for _, o := range options {
		o(&opt)
	}
	// 検索オプションの内容からクエリを作成
	query := bson.M{}
	if len(opt.keyword) > 0 {
		keywords := strings.Split(opt.keyword, " ")
		for _, keyword := range keywords {
			regex := bson.M{"$regex": bson.RegEx{Pattern: `.*` + keyword + `.*`, Options: "im"}}
			query = bson.M{
				"$and": []interface{}{
					query,
					bson.M{
						"$or": []interface{}{
							bson.M{"title": regex},
							bson.M{"body": regex},
							bson.M{"place.name": regex},
							bson.M{"host.name": regex},
						},
					},
				},
			}
		}
	}
	if len(opt.userID) != 0 {
		if len(opt.keyword) > 0 {
			query = bson.M{
				"$and": []interface{}{
					query,
					bson.M{"host.id": opt.userID},
				},
			}
		} else {
			query = bson.M{"host.id": opt.userID}
		}
	}
	err = this.Find(query).Select(EVENT_TINY_SELECTOR).Skip(offset).Limit(limit).All(&events)
	return
}
