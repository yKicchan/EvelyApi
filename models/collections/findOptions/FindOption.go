package findOptions

// 検索時に使用するオプション
type findOption struct {
	// 検索上限
	limit int
	// 除外件数(0があり得るのでセット判定のためにポインタ型)
	offset *int
	// カテゴリの絞り込み
	categorys []string
}

// 汎用的な検索オプション
type FindOption interface {
	SetLimit(int)
	SetOffset(int)
	SetCategorys([]string)
	GetLimit() int
	GetOffset() int
	GetCategorys() []string
	IsLimitSet() bool
	IsOffsetSet() bool
	IsCategorysSet() bool
}

/**
 * 検索件数の上限を設定する
 * @param limit 上限件数
 */
func (this *findOption) SetLimit(limit int) {
	if limit > 0 {
		this.limit = limit
	}
}

/**
 * 検索時のスキップする除外件数を設定する
 * @param offset 除外件数
 */
func (this *findOption) SetOffset(offset int) {
	if offset >= 0 {
		this.offset = &offset
	}
}

/**
 * 検索時にカテゴリで絞り込む
 * @param categoy カテゴリ
 */
func (this *findOption) SetCategorys(categorys []string) {
	if categorys != nil {
		this.categorys = categorys
	}
}

/**
 * 検索上限件数を返す
 * @return int 件数
 */
func (this *findOption) GetLimit() int { return this.limit }

/**
 * 検索時にスキップする除外件数を返す
 * @return int 除外件数
 */
func (this *findOption) GetOffset() int { return *this.offset }

/**
 * 検索時に絞り込む
 * @return string カテゴリ
 */
func (this *findOption) GetCategorys() []string { return this.categorys }

/**
 * 検索件数に上限が設定されているかを判定する
 * @return bool
 */
func (this *findOption) IsLimitSet() bool { return this.limit > 0 }

/**
 * 検索時にスキップする除外件数が設定されているかを判定する
 * @return bool
 */
func (this *findOption) IsOffsetSet() bool { return this.offset != nil }

func (this *findOption) IsCategorysSet() bool { return this.categorys != nil }