package collection

import (
	"labix.org/v2/mgo/bson"
)

// ドキュメントを特定するためのキー
type Keys map[string]interface{}

/**
 * 検索クエリの形に変換する
 * @return query MongoDB用のクエリに変換したマップ
 */
func (this Keys) ToQuery() (query bson.M) {
	for key, val := range this {
		query = bson.M{
			"$and": []interface{}{
				query,
				bson.M{key: val},
			},
		}
	}
	return query
}
