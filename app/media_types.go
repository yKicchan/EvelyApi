// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": Application Media Types
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

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
	// 作成日時
	CreatedAt time.Time `form:"createdAt" json:"createdAt" xml:"createdAt"`
	Host      *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// 連絡先メールアドレス
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// 通知範囲(m)
	NoticeRange int `form:"noticeRange" json:"noticeRange" xml:"noticeRange"`
	// 開催中かどうか
	OpenFlg bool `form:"openFlg" json:"openFlg" xml:"openFlg"`
	// イベントの開催予定一覧
	Plans []*Plan `form:"plans" json:"plans" xml:"plans"`
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

// Validate validates the Event media type instance.
func (mt *Event) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "body"))
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
	if mt.Plans == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "plans"))
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
	for _, e := range mt.Plans {
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
	Host *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// 通知範囲(m)
	NoticeRange int `form:"noticeRange" json:"noticeRange" xml:"noticeRange"`
	// 開催中かどうか
	OpenFlg bool `form:"openFlg" json:"openFlg" xml:"openFlg"`
	// イベントの開催予定一覧
	Plans []*Plan `form:"plans" json:"plans" xml:"plans"`
	// 公開範囲
	Scope string `form:"scope" json:"scope" xml:"scope"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
	// 最終更新日時
	UpdateDate time.Time `form:"updateDate" json:"updateDate" xml:"updateDate"`
}

// Validate validates the EventTiny media type instance.
func (mt *EventTiny) Validate() (err error) {
	if mt.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "id"))
	}
	if mt.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "title"))
	}
	if mt.Host == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "host"))
	}
	if mt.Plans == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "plans"))
	}

	if mt.Scope == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "scope"))
	}

	if mt.Host != nil {
		if err2 := mt.Host.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.NoticeRange < 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.noticeRange`, mt.NoticeRange, 100, true))
	}
	if mt.NoticeRange > 5000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`response.noticeRange`, mt.NoticeRange, 5000, false))
	}
	for _, e := range mt.Plans {
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

// アップロード済みのファイルのパス (default view)
//
// Identifier: application/vnd.file_path+json; view=default
type FilePath struct {
	// ファイルのパス
	Path string `form:"path" json:"path" xml:"path"`
}

// Validate validates the FilePath media type instance.
func (mt *FilePath) Validate() (err error) {
	if mt.Path == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "path"))
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
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// メールアドレス
	Mail *Mail `form:"mail" json:"mail" xml:"mail"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
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
	if mt.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mail"))
	}
	if mt.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tel"))
	}

	if utf8.RuneCountInString(mt.ID) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 4, true))
	}
	if utf8.RuneCountInString(mt.ID) > 15 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 15, false))
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
	return
}

// ユーザー情報 (tiny view)
//
// Identifier: application/vnd.user+json; view=tiny
type UserTiny struct {
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
	if utf8.RuneCountInString(mt.ID) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 4, true))
	}
	if utf8.RuneCountInString(mt.ID) > 15 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.id`, mt.ID, utf8.RuneCountInString(mt.ID), 15, false))
	}
	if utf8.RuneCountInString(mt.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 1, true))
	}
	if utf8.RuneCountInString(mt.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`response.name`, mt.Name, utf8.RuneCountInString(mt.Name), 50, false))
	}
	return
}
