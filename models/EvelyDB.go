package models

import (
	. "EvelyApi/config"
	. "EvelyApi/models/collections"
	. "EvelyApi/models/collections/findOptions"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

/**
 * Evelyのデーターベースを操作するためのオブジェクト
 */
type EvelyDB struct {
	*mgo.Database
	Events  *EventsCollection
	Users   *UsersCollection
	Reviews *ReviewsCollection
}

func NewEvelyDB(db *mgo.Database) *EvelyDB {
	return &EvelyDB{
		db,
		NewEventsCollection(db.C(EVENTS_COLLECTION)),
		NewUsersCollection(db.C(USERS_COLLECTION)),
		NewReviewsCollection(db.C(REVIEWS_COLLECTION)),
	}
}

/**
 * イベントを位置情報から検索する
 * 現在地からの距離あり
 * @param  opt 検索オプション
 * @return results 検索結果
 * @return err     エラー
 */
func (this *EvelyDB) FindEventsByLocation(opt LocationOption) (results []*Result, err error) {
	// 位置情報検索クエリ生成
	lat, lng, r := opt.GetLocation()
	// リザルトを受け取る時
	command := bson.M{
		"geoNear": EVENTS_COLLECTION,
		"near": bson.M{
			"type": "Point",
			"coordinates": []float64{lng, lat},
		},
		"spherical": true,
		"maxDistance": r,
	}
	if opt.IsLimitSet() {
		command["limit"] = opt.GetLimit()
	}
	if opt.IsCategorysSet() {
		categorys, ope := opt.GetCategorys()
		command["query"] = bson.M{"categorys": bson.M{ope: categorys}}
	}
	res := &RunResult{}
	err = this.Run(command, res)
	return res.Results, err
}
