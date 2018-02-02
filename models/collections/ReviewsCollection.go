package collections

import (
	. "EvelyApi/models/documents"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
)

/**
 * Reviewsコレクションを操作する
 */
type ReviewsCollection struct {
	*mgo.Collection
}

func NewReviewsCollection(c *mgo.Collection) *ReviewsCollection {
	return &ReviewsCollection{c}
}

/**
 * レビューを保存(更新)する
 * @param  review レビュー情報
 * @param  keys   レビューを特定するキー
 * @return error  保存時に発生したエラー
 */
func (this *ReviewsCollection) Save(review *ReviewModel, keys Keys) error {
	update := bson.M{"$set": review}
	_, err := this.Upsert(keys.ToQuery(), update)
	return err
}

/**
 * 条件にあうレビューを全て更新する
 * @param event レビュー情報
 * @param keys  検索条件
 */
func (this *ReviewsCollection) SaveAll(review *ReviewModel, query bson.M) error {
	update := bson.M{"$set": review}
	info, err := this.UpdateAll(query, update)
	log.Printf("[Events] SaveAll info: %v", info)
	return err
}

/**
 * レビューを検索する
 * @param  keys レビューを特定するキー
 * @return user レビューの情報
 * @return err  エラー
 */
func (this *ReviewsCollection) FindOne(keys Keys) (r *ReviewModel, err error) {
	err = this.Find(keys.ToQuery()).One(&r)
	return
}

/**
 * レビューを削除する
 * @param  keys  レビューを特定するキー
 * @return error エラー
 */
func (this *ReviewsCollection) Delete(keys Keys) error {
	return this.Remove(keys.ToQuery())
}
