package collections

import (
	. "EvelyApi/models/collections/findOptions"
	. "EvelyApi/models/documents"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strings"
)

/**
 * Eventsコレクションを操作するためのオブジェクト
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
 * イベントをクエリとオプションから検索する
 * @param  opt    検索オプション
 * @param  query  クエリ
 * @return events イベント配列
 * @return err    検索時のエラー
 */
func (this *EventsCollection) findEvents(opt FindOption, query bson.M) (events []*EventModel, err error) {
	if opt.IsCategorysSet() {
		categorys, ope := opt.GetCategorys()
		query = bson.M{
			"$and": []interface{}{
				query,
				bson.M{"categorys": bson.M{ope: categorys}},
			},
		}
	}
	q := this.Find(query).Select(EVENT_TINY_SELECTOR)
	if opt.IsOffsetSet() {
		q = q.Skip(opt.GetOffset())
	}
	if opt.IsLimitSet() {
		q = q.Limit(opt.GetLimit())
	}
	err = q.All(&events)
	return
}

/**
 * イベントを検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEvents(opt *FindEventsOption) (events []*EventModel, err error) {
	if opt.IsKeywordSet() {
		return this.findEventsByKeyword(opt)
	}
	if opt.IsLocationSet() {
		return this.findEventsByLocation(opt)
	}
	if opt.IsHostIDSet() {
		return this.findEventsByHostID(opt)
	}
	return this.findEvents(opt, nil)
}

/**
 * イベントをキーワードから検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) findEventsByKeyword(opt KeywordOption) (events []*EventModel, err error) {
	// キーワード検索クエリ生成
	var query bson.M
	keywords := strings.Split(opt.GetKeyword(), " ")
	for _, keyword := range keywords {
		regex := bson.M{"$regex": bson.RegEx{Pattern: `.*` + keyword + `.*`, Options: "im"}}
		query = bson.M{
			"$and": []interface{}{
				query,
				bson.M{
					"$or": []interface{}{
						bson.M{"title": regex},
						bson.M{"body": regex},
						bson.M{"schedules": bson.M{"$elemMatch": bson.M{"location.name": regex}}},
						bson.M{"host.name": regex},
					},
				},
			},
		}
	}
	return this.findEvents(opt, query)
}

/**
 * イベントを位置情報から検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) findEventsByLocation(opt LocationOption) (events []*EventModel, err error) {
	// 位置情報検索クエリ生成
	lat, lng, r := opt.GetLocation()
	query := bson.M{
		"schedules.location.lng_lat": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{lng, lat},
				},
				"$maxDistance": r,
			},
		},
	}
	return this.findEvents(opt, query)
}

/**
 * イベントを作成者から検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) findEventsByHostID(opt HostIDOption) (events []*EventModel, err error) {
	// 作成者からの検索クエリ生成
	query := bson.M{"host.id": opt.GetHostID()}
	return this.findEvents(opt, query)
}
