package collection

// 検索オプション
type findOptions struct {
	limit       int
	offset      int
	keyword     string
	lat         float64
	lng         float64
	r           float64
	instanceID string
}

// 検索オプションを設定するため関数化
type FindOptions func(*findOptions)

/**
 * 検索件数の上限を設定する
 * @param  limit 検索件数
 * @return       クロージャ
 */
func WithLimit(limit int) FindOptions {
	return func(ops *findOptions) {
		ops.limit = limit
	}
}

/**
 * 検索時、スキップする件数を設定する
 * @param  offset 除外件数
 * @return        クロージャ
 */
func WithOffset(offset int) FindOptions {
	return func(ops *findOptions) {
		ops.offset = offset
	}
}

/**
 * キーワードを検索オプションに設定する
 * @param  keyword キーワード
 * @return         クロージャ
 */
func WithKeyword(keyword string) FindOptions {
	return func(ops *findOptions) {
		ops.keyword = keyword
	}
}

/**
 * 位置情報を検索オプションに設定する
 * @param  lat 緯度
 * @param  lng 経度
 * @param  r   検索範囲(度)
 * @return     クロージャ
 */
func WithLocation(lat, lng, r float64) FindOptions {
	return func(ops *findOptions) {
		ops.lat = lat
		ops.lng = lng
		ops.r = r
	}
}

/**
 * インスタンスIDを検索オプションに設定する
 * @param  instanceID インスタンスID
 * @return             クロージャ
 */
func WithInstanceID(instanceID string) FindOptions {
	return func(ops *findOptions) {
		ops.instanceID = instanceID
	}
}
