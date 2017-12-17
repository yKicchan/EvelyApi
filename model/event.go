package model

import (
    . "EvelyApi/config"
    "labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"strconv"
	"time"
)

// イベントのリソースに対するDBオブジェクト
type EventDB struct {
	*mgo.Database
}

// イベントの検索オプション
type getEventsOptions struct {
	keyword string
	userID  string
}

// イベントの検索オプションを関数化
type GetEventsOption func(*getEventsOptions)

/**
 * ユーザーIDを検索オプションに設定するクロージャを返す
 * @param  userID ユーザーID
 * @return        クロージャ
 */
func WithUserID(userID string) GetEventsOption {
	return func(ops *getEventsOptions) {
		ops.userID = userID
	}
}

/**
 * キーワードを検索オプションに設定するクロージャを返す
 * @param  keyword キーワード
 * @return         クロージャ
 */
func WithKeyword(keyword string) GetEventsOption {
	return func(ops *getEventsOptions) {
		ops.keyword = keyword
	}
}

func NewEventDB(db *mgo.Database) *EventDB {
	return &EventDB{db}
}

/**
 * イベントを新しく作成するためのイベントIDを生成し返す
 * @param  userID  作成者のユーザーID
 * @return eventID 生成したイベントのID
 * @return err     作成時に発生したエラー
 */
func (db *EventDB) NewEvent(userID string, upcomingDate time.Time) (eventID string, err error) {
	y, m, d := upcomingDate.Date()
	date := strconv.Itoa(y) + strconv.Itoa(int(m)) + strconv.Itoa(d)
	query := bson.M{
		"$and": []interface{}{
			bson.M{"id": bson.M{"$regex": bson.RegEx{Pattern: date + `-[0-9]+`}}},
			bson.M{"host.id": userID},
		},
	}
	n, _ := db.C("events").Find(query).Count()
	eventID = date + "-" + strconv.Itoa(n+1)
	event := &EventModel{
		ID:   eventID,
		Host: Host{ID: userID},
	}
	err = db.C("events").Insert(event)
	return
}

/**
 * イベント情報をパラメーターから検索し、複数件返す
 * @param  limit   検索件数
 * @param  offset  除外件数
 * @param  options 検索オプション
 * @return events  複数のイベント情報
 * @return err     検索時に発生したエラー
 */
func (db *EventDB) GetEvents(limit, offset int, options ...GetEventsOption) (events []*EventModel, err error) {
	// 検索オプションを取得
	opt := getEventsOptions{}
	for _, o := range options {
		o(&opt)
	}
	// 検索オプションの内容からクエリを作成
	query := bson.M{}
	if len(opt.keyword) > 0 {
		regex := bson.M{"$regex": bson.RegEx{Pattern: `.*` + opt.keyword + `.*`, Options: "m"}}
		query = bson.M{
			"$or": []interface{}{
				bson.M{"title": regex},
				bson.M{"body": regex},
				bson.M{"place": regex},
				bson.M{"host.name": regex},
			},
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
	err = db.C("events").Find(query).Select(EVENT_TINY_SELECTOR).Skip(offset).Limit(limit).All(&events)
	return
}

/**
 * イベントの詳しい情報を返す
 * @param  userID  主催者のユーザーID
 * @param  eventID イベントID
 * @return event   イベント情報
 * @return err     検索時に発生したエラー
 */
func (db *EventDB) GetEvent(userID, eventID string) (event *EventModel, err error) {
	query := bson.M{
		"$and": []interface{}{
			bson.M{"id": eventID},
			bson.M{"host.id": userID},
		},
	}
	err = db.C("events").Find(query).Select(FULL_SELECTOR).One(&event)
	return
}

/**
 * イベントを保存する
 * @param  event イベント情報
 * @return error 保存時に発生したエラー
 */
func (db *EventDB) SaveEvent(event *EventModel) error {
	selector := bson.M{
		"$and": []interface{}{
			bson.M{"id": event.ID},
			bson.M{"host.id": event.Host.ID},
		},
	}
	update := bson.M{"$set": event}
	return db.C("events").Update(selector, update)
}

/**
 * イベントを削除する
 * @param  userID  ユーザーID
 * @param  eventID イベントID
 * @return error   削除時に発生したエラー
 */
func (db *EventDB) DeleteEvent(userID, eventID string) error {
	selector := bson.M{
		"$and": []interface{}{
			bson.M{"id": eventID},
			bson.M{"host.id": userID},
		},
	}
	return db.C("events").Remove(selector)
}
