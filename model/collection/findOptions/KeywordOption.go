package findOptions

// キーワード検索のオプション
type KeywordOption interface {
    FindOptions
    SetKeyword(string)
    GetKeyword() string
    IsKeywordSet() bool
}
