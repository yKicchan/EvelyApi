package collection

import (
	. "EvelyApi/config"
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
func (this *EventsCollection) Save(event *EventModel, keys Keys) error {
	update := bson.M{"$set": event}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * イベントを検索する
 * @param  keys  イベントを特定するキー
 * @return err   エラー
 */
func (this *EventsCollection) FindOne(keys Keys) (e *EventModel, err error) {
	err = this.Find(keys.ToQuery()).One(&e)
	return
}

/**
 * イベントを削除する
 * @param  keys  イベントを特定するキー
 * @return error エラー
 */
func (this *EventsCollection) Delete(keys Keys) error {
	return this.Remove(keys.ToQuery())
}

/**
 * イベント情報をパラメーターから検索し、複数件返す
 * @param  options 検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEvents(options ...FindOptions) (events []*EventModel, err error) {
	// 検索オプションを取得、設定
	opt := findOptions{}
	for _, o := range options {
		o(&opt)
	}
	// 検索オプションの内容からクエリを作成
	var query bson.M
	// キーワード検索
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
							bson.M{"plans": bson.M{"$elemMatch": bson.M{"location.name": regex}}},
							bson.M{"host.name": regex},
						},
					},
				},
			}
		}
	}
	// 位置情報検索
	if opt.r > 0 {
		query = bson.M{
			"plans.location.lng_lat": bson.M{
				"$nearSphere": bson.M{
					"$geometry": bson.M{
						"type":        "Point",
						"coordinates": []float64{opt.lng, opt.lat},
					},
					"$maxDistance": (float64(opt.r) * DEGREE_PER_METER),
				},
			},
		}
	}

	// 検索条件からイベントを検索
	q := this.Find(query).Select(EVENT_TINY_SELECTOR)

	// 除外件数を指定
	if opt.offset > 0 {
		q = q.Skip(opt.offset)
	}
	// 検索件数の上限を指定
	if opt.limit > 0 {
		q = q.Limit(opt.limit)
	}
	// 結果を返す
	err = q.All(&events)
	return
}
