package collection

import (
    . "EvelyApi/model/document"
    "labix.org/v2/mgo/bson"
)

/**
 * Evelyで使用するコレクションのインターフェース
 */
type EvelyCollection interface {

	/**
	 * ドキュメントを保存(更新)する
	 * @param  model ドキュメント
	 * @param  keys  ドキュメントを特定するためのキー
	 * @return error クエリ実行時のエラー内容
	 */
	Save(model EvelyModel, keys Keys) error

	/**
	 * ドキュメントを検索する
	 * @param  keys  ドキュメントを特定するキー
	 * @return model 検索にヒットしたドキュメント
	 * @return err   クエリ実行時のエラー内容
	 */
	FindDoc(keys Keys) (EvelyModel, error)

	/**
	 * ドキュメントを削除する
	 * @param  keys  ドキュメントを特定するキー
	 * @return error クエリ実行時のエラー内容
	 */
	Delete(keys Keys) error
}

// ドキュメントを特定するためのキー
type Keys map[string]interface{}

/**
 * 検索クエリの形に変換する
 * @return query MongoDB用のクエリに変換したマップ
 */
func (this Keys) ToQuery() (query bson.M) {
    for key, val := range this {
        if key == "_id" {
            val = bson.ObjectIdHex(val.(string))
        }
        query = bson.M{
            "$and": []interface{}{
                query,
                bson.M{key: val},
            },
        }
    }
    return query
}
