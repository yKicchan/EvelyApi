package findOptions

// イベント検索で使用する検索オプション
type FindEventsOption struct {
	findOption
	keywordOption
	hostIDOption
	locationOption
}

func NewFindEventsOption() *FindEventsOption {
	return &FindEventsOption{}
}
