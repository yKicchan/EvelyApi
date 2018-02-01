// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "EvelyApi": Application Media Types
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.1

package app

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// メールアドレス (default view)
//
// Identifier: application/vnd.email+json; view=default
type Email struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
}

// Validate validates the Email media type instance.
func (mt *Email) Validate() (err error) {
	if mt.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.email`, mt.Email, goa.FormatEmail, err2))
	}
	return
}

// イベント情報 (default view)
//
// Identifier: application/vnd.event+json; view=default
type Event struct {
	// イベントの詳細
	Body string `form:"body" json:"body" xml:"body"`
	// カテゴリ
	Categorys []string `form:"categorys" json:"categorys" xml:"categorys"`
	// 作成日時
	CreatedAt time.Time `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// 添付資料へのURL
	Files []string  `form:"files" json:"files" xml:"files"`
	Host  *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// ヘッダー画像のURL
	Image string `form:"image" json:"image" xml:"image"`
	// レビューの有無
	IsReviewed bool `form:"isReviewed" json:"isReviewed" xml:"isReviewed"`
	// 連絡先メールアドレス
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// イベントの開催予定一覧
	Schedules []*Schedule `form:"schedules" json:"schedules" xml:"schedules"`
	// 連絡先電話番号
	Tel string `form:"tel" json:"tel" xml:"tel"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
	// 最終更新日時
	UpdateDate time.Time `form:"updateDate" json:"updateDate" xml:"updateDate"`
	// URL
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the Event media type instance.
func (mt *Event) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if mt.Files == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "files"))
	}
	if mt.Categorys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "categorys"))
	}
	if mt.Host == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "host"))
	}
	if mt.Mail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mail"))
	}
	if mt.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tel"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Schedules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "schedules"))
	}

	if utf8.RuneCountInString(mt.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1, true))
	}
	if utf8.RuneCountInString(mt.Body) > 1000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1000, false))
	}
	if len(mt.Categorys) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 1, true))
	}
	if len(mt.Categorys) > 9 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 9, false))
	}
	if mt.Host != nil {
		if err2 := mt.Host.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Mail); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.mail`, mt.Mail, goa.FormatEmail, err2))
	}
	if len(mt.Schedules) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.schedules`, mt.Schedules, len(mt.Schedules), 1, true))
	}
	for _, e := range mt.Schedules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(mt.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 1, true))
	}
	if utf8.RuneCountInString(mt.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 30, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.URL); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.url`, mt.URL, goa.FormatURI, err2))
	}
	return
}

// イベント情報 (full view)
//
// Identifier: application/vnd.event+json; view=full
type EventFull struct {
	// イベントの詳細
	Body string `form:"body" json:"body" xml:"body"`
	// カテゴリ
	Categorys []string `form:"categorys" json:"categorys" xml:"categorys"`
	// 作成日時
	CreatedAt time.Time `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// 添付資料へのURL
	Files []string  `form:"files" json:"files" xml:"files"`
	Host  *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// ヘッダー画像のURL
	Image string `form:"image" json:"image" xml:"image"`
	// レビューの有無
	IsReviewed bool `form:"isReviewed" json:"isReviewed" xml:"isReviewed"`
	// 連絡先メールアドレス
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// 通知範囲(m)
	NoticeRange int `form:"noticeRange" json:"noticeRange" xml:"noticeRange"`
	// 開催中かどうか
	OpenFlg bool `form:"openFlg" json:"openFlg" xml:"openFlg"`
	// イベントの開催予定一覧
	Schedules []*Schedule `form:"schedules" json:"schedules" xml:"schedules"`
	// 公開範囲
	Scope string `form:"scope" json:"scope" xml:"scope"`
	// 連絡先電話番号
	Tel string `form:"tel" json:"tel" xml:"tel"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
	// 最終更新日時
	UpdateDate time.Time `form:"updateDate" json:"updateDate" xml:"updateDate"`
	// URL
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the EventFull media type instance.
func (mt *EventFull) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if mt.Files == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "files"))
	}
	if mt.Categorys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "categorys"))
	}
	if mt.Host == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "host"))
	}
	if mt.Mail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mail"))
	}
	if mt.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tel"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Schedules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "schedules"))
	}

	if mt.Scope == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "scope"))
	}

	if utf8.RuneCountInString(mt.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1, true))
	}
	if utf8.RuneCountInString(mt.Body) > 1000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1000, false))
	}
	if len(mt.Categorys) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 1, true))
	}
	if len(mt.Categorys) > 9 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 9, false))
	}
	if mt.Host != nil {
		if err2 := mt.Host.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, mt.Mail); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.mail`, mt.Mail, goa.FormatEmail, err2))
	}
	if mt.NoticeRange < 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.noticeRange`, mt.NoticeRange, 100, true))
	}
	if mt.NoticeRange > 5000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.noticeRange`, mt.NoticeRange, 5000, false))
	}
	if len(mt.Schedules) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.schedules`, mt.Schedules, len(mt.Schedules), 1, true))
	}
	for _, e := range mt.Schedules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if !(mt.Scope == "public" || mt.Scope == "private") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.scope`, mt.Scope, []interface{}{"public", "private"}))
	}
	if utf8.RuneCountInString(mt.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 1, true))
	}
	if utf8.RuneCountInString(mt.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 30, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, mt.URL); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`response.url`, mt.URL, goa.FormatURI, err2))
	}
	return
}

// イベント情報 (tiny view)
//
// Identifier: application/vnd.event+json; view=tiny
type EventTiny struct {
	// カテゴリ
	Categorys []string  `form:"categorys" json:"categorys" xml:"categorys"`
	Host      *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// ヘッダー画像のURL
	Image string `form:"image" json:"image" xml:"image"`
	// レビューの有無
	IsReviewed bool `form:"isReviewed" json:"isReviewed" xml:"isReviewed"`
	// イベントの開催予定一覧
	Schedules []*Schedule `form:"schedules" json:"schedules" xml:"schedules"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the EventTiny media type instance.
func (mt *EventTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Image == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "image"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Categorys == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "categorys"))
	}
	if mt.Host == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "host"))
	}
	if mt.Schedules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "schedules"))
	}

	if len(mt.Categorys) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 1, true))
	}
	if len(mt.Categorys) > 9 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.categorys`, mt.Categorys, len(mt.Categorys), 9, false))
	}
	if mt.Host != nil {
		if err2 := mt.Host.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if len(mt.Schedules) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.schedules`, mt.Schedules, len(mt.Schedules), 1, true))
	}
	for _, e := range mt.Schedules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if utf8.RuneCountInString(mt.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 1, true))
	}
	if utf8.RuneCountInString(mt.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 30, false))
	}
	return
}

// EventCollection is the media type for an array of Event (default view)
//
// Identifier: application/vnd.event+json; type=collection; view=default
type EventCollection []*Event

// Validate validates the EventCollection media type instance.
func (mt EventCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// EventCollection is the media type for an array of Event (full view)
//
// Identifier: application/vnd.event+json; type=collection; view=full
type EventFullCollection []*EventFull

// Validate validates the EventFullCollection media type instance.
func (mt EventFullCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// EventCollection is the media type for an array of Event (tiny view)
//
// Identifier: application/vnd.event+json; type=collection; view=tiny
type EventTinyCollection []*EventTiny

// Validate validates the EventTinyCollection media type instance.
func (mt EventTinyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// 近くのイベント (default view)
//
// Identifier: application/vnd.nearby+json; view=default
type Nearby struct {
	// イベントまでの距離(m)
	Distance int        `form:"distance" json:"distance" xml:"distance"`
	Event    *EventTiny `form:"event" json:"event" xml:"event"`
}

// Validate validates the Nearby media type instance.
func (mt *Nearby) Validate() (err error) {
	if mt.Event == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "event"))
	}

	if mt.Event != nil {
		if err2 := mt.Event.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// NearbyCollection is the media type for an array of Nearby (default view)
//
// Identifier: application/vnd.nearby+json; type=collection; view=default
type NearbyCollection []*Nearby

// Validate validates the NearbyCollection media type instance.
func (mt NearbyCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// レビュー (default view)
//
// Identifier: application/vnd.review+json; view=default
type Review struct {
	// レビューの内容
	Body string `form:"body" json:"body" xml:"body"`
	// レビュー画像などのURL
	Files []string `form:"files" json:"files" xml:"files"`
	// レビューID
	ID string `form:"id" json:"id" xml:"id"`
	// レート評価
	Rate int `form:"rate" json:"rate" xml:"rate"`
	// レビューした日
	ReviewedAt time.Time `form:"reviewedAt" json:"reviewedAt" xml:"reviewedAt"`
	Reviewer   *UserTiny `form:"reviewer" json:"reviewer" xml:"reviewer"`
	// レビューのタイトル
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the Review media type instance.
func (mt *Review) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}

	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
	}
	if mt.Files == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "files"))
	}
	if mt.Reviewer == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "reviewer"))
	}

	if utf8.RuneCountInString(mt.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 1, true))
	}
	if utf8.RuneCountInString(mt.Body) > 5000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.body`, mt.Body, utf8.RuneCountInString(mt.Body), 5000, false))
	}
	if !(mt.Rate == 1 || mt.Rate == 2 || mt.Rate == 3 || mt.Rate == 4 || mt.Rate == 5) {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`response.rate`, mt.Rate, []interface{}{1, 2, 3, 4, 5}))
	}
	if mt.Reviewer != nil {
		if err2 := mt.Reviewer.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(mt.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 1, true))
	}
	if utf8.RuneCountInString(mt.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.title`, mt.Title, utf8.RuneCountInString(mt.Title), 30, false))
	}
	return
}

// ReviewCollection is the media type for an array of Review (default view)
//
// Identifier: application/vnd.review+json; type=collection; view=default
type ReviewCollection []*Review

// Validate validates the ReviewCollection media type instance.
func (mt ReviewCollection) Validate() (err error) {
	for _, e := range mt {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	return
}

// アクセストークン (default view)
//
// Identifier: application/vnd.token+json; view=default
type Token struct {
	// アクセストークン
	Token string `form:"token" json:"token" xml:"token"`
}

// Validate validates the Token media type instance.
func (mt *Token) Validate() (err error) {
	if mt.Token == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "token"))
	}
	return
}

// ユーザー情報 (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
	// 作成日時
	CreatedAt time.Time `form:"createdAt" json:"createdAt" xml:"createdAt"`
	// アイコン画像のURL
	Icon string `form:"icon" json:"icon" xml:"icon"`
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// メールアドレス
	Mail *Mail `form:"mail" json:"mail" xml:"mail"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
	// ピンしているイベントのID配列
	Pins []string `form:"pins" json:"pins" xml:"pins"`
	// 通知を許可するカテゴリの配列
	Preferences []string `form:"preferences" json:"preferences" xml:"preferences"`
	// 電話番号
	Tel string `form:"tel" json:"tel" xml:"tel"`
}

// Validate validates the User media type instance.
func (mt *User) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Icon == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "icon"))
	}
	if mt.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mail"))
	}
	if mt.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tel"))
	}
	if mt.Pins == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "pins"))
	}
	if mt.Preferences == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "preferences"))
	}

	if ok := goa.ValidatePattern(`^[a-zA-Z0-9_]{4,15}$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[a-zA-Z0-9_]{4,15}$`))
	}
	if mt.Mail != nil {
		if err2 := mt.Mail.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 50, false))
	}
	if len(mt.Preferences) > 9 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.preferences`, mt.Preferences, len(mt.Preferences), 9, false))
	}
	return
}

// ユーザー情報 (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
	// アイコン画像のURL
	Icon string `form:"icon" json:"icon" xml:"icon"`
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the UserTiny media type instance.
func (mt *UserTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "name"))
	}
	if mt.Icon == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "icon"))
	}
	if ok := goa.ValidatePattern(`^[a-zA-Z0-9_]{4,15}$`, mt.ID); !ok {
		err = goa.MergeErrors(err, goa.InvalidPatternError(`response.id`, mt.ID, `^[a-zA-Z0-9_]{4,15}$`))
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 50, false))
	}
	return
}
