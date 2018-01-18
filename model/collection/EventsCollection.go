package collection

import (
	. "EvelyApi/model/collection/findOptions"
	. "EvelyApi/model/document"
	"errors"
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
 * イベントを検索する
 * @param  opt 検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEvents(opt FindOption) (events []*EventModel, err error) {
    // 検索条件からイベントを検索
	q := this.Find(nil).Select(EVENT_TINY_SELECTOR)
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
 * イベントをキーワードから検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEventsByKeyword(opt KeywordOption) (events []*EventModel, err error) {
	// 設定されているかチェック
	if !opt.IsKeywordSet() {
		return nil, errors.New("キーワードが設定されていません")
	}
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

	// 検索条件からイベントを検索
	q := this.Find(query).Select(EVENT_TINY_SELECTOR)

	// 除外件数を指定
	if opt.IsOffsetSet() {
		q = q.Skip(opt.GetOffset())
	}
	// 検索件数の上限を指定
	if opt.IsLimitSet() {
		q = q.Limit(opt.GetLimit())
	}
	// 結果を返す
	err = q.All(&events)
	return
}

/**
 * イベントを位置情報から検索する
 * @param  opt     検索時のオプション
 * @return events  検索にヒットした複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (this *EventsCollection) FindEventsByLocation(opt LocationOption) (events []*EventModel, err error) {
	// 検索オプションが設定されているかチェック
	if !opt.IsLocationSet() {
		return nil, errors.New("位置情報が設定されていません")
	}
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

	// 検索条件からイベントを検索
	q := this.Find(query).Select(EVENT_TINY_SELECTOR)

	// 除外件数を指定
	if opt.IsOffsetSet() {
		q = q.Skip(opt.GetOffset())
	}
	// 検索件数の上限を指定
	if opt.IsLimitSet() {
		q = q.Limit(opt.GetLimit())
	}
	// 結果を返す
	err = q.All(&events)
	return
}
