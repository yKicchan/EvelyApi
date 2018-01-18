package findOptions

// キーワード検索のオプション
type KeywordOption interface {
	FindOption
	SetKeyword(string)
	GetKeyword() string
	IsKeywordSet() bool
}
