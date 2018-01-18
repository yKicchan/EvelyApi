package findOptions

// キーワード検索のオプション
type KeywordOption interface {
    FindOption
	SetKeyword(string)
	GetKeyword() string
	IsKeywordSet() bool
}

type keywordOption struct {
	keyword string
}

/**
 * キーワードを検索オプションに設定する
 * @param  keyword キーワード
 */
func (this *keywordOption) SetKeyword(keyword string) {
	if keyword != "" {
		this.keyword = keyword
	}
}

/**
 * 設定されたキーワードを返す
 * @return キーワード
 */
func (this *keywordOption) GetKeyword() string { return this.keyword }

/**
 * キーワードが有効かを判定する
 * @return bool
 */
func (this *keywordOption) IsKeywordSet() bool { return this.keyword != "" }
