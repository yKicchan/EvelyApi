// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": Application Media Types
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"net/http"
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

// DecodeEmail decodes the Email instance encoded in resp body.
func (c *Client) DecodeEmail(resp *http.Response) (*Email, error) {
	var decoded Email
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// イベント情報 (default view)
//
// Identifier: application/vnd.event+json; view=default
type Event struct {
	// イベントの詳細
	Body string    `form:"body" json:"body" xml:"body"`
	Host *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// 連絡先メールアドレス
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// 開催場所
	Place *Location `form:"place" json:"place" xml:"place"`
	// 連絡先電話番号
	Tel string `form:"tel" json:"tel" xml:"tel"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
	// 開催予定日
	UpcomingDate *UpcomingDate `form:"upcomingDate" json:"upcomingDate" xml:"upcomingDate"`
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
	if mt.Place == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "place"))
	}

	if mt.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "upcomingDate"))
	}
	if mt.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "url"))
	}
	if mt.Mail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "mail"))
	}
	if mt.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "tel"))
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
	if mt.Place != nil {
		if err2 := mt.Place.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
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

// イベント情報 (tiny view)
//
// Identifier: application/vnd.event+json; view=tiny
type EventTiny struct {
	Host *UserTiny `form:"host" json:"host" xml:"host"`
	// イベントID
	ID string `form:"id" json:"id" xml:"id"`
	// 開催場所
	Place *Location `form:"place" json:"place" xml:"place"`
	// イベントの名前
	Title string `form:"title" json:"title" xml:"title"`
	// 開催予定日
	UpcomingDate *UpcomingDate `form:"upcomingDate" json:"upcomingDate" xml:"upcomingDate"`
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
	if mt.Place == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "place"))
	}
	if mt.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`response`, "upcomingDate"))
	}
	if mt.Host != nil {
		if err2 := mt.Host.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if mt.Place != nil {
		if err2 := mt.Place.Validate(); err2 != nil {
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

// DecodeEvent decodes the Event instance encoded in resp body.
func (c *Client) DecodeEvent(resp *http.Response) (*Event, error) {
	var decoded Event
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeEventTiny decodes the EventTiny instance encoded in resp body.
func (c *Client) DecodeEventTiny(resp *http.Response) (*EventTiny, error) {
	var decoded EventTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
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

// DecodeEventCollection decodes the EventCollection instance encoded in resp body.
func (c *Client) DecodeEventCollection(resp *http.Response) (EventCollection, error) {
	var decoded EventCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
}

// DecodeEventTinyCollection decodes the EventTinyCollection instance encoded in resp body.
func (c *Client) DecodeEventTinyCollection(resp *http.Response) (EventTinyCollection, error) {
	var decoded EventTinyCollection
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return decoded, err
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

// DecodeToken decodes the Token instance encoded in resp body.
func (c *Client) DecodeToken(resp *http.Response) (*Token, error) {
	var decoded Token
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// ユーザー情報 (default view)
//
// Identifier: application/vnd.user+json; view=default
type User struct {
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

// DecodeUser decodes the User instance encoded in resp body.
func (c *Client) DecodeUser(resp *http.Response) (*User, error) {
	var decoded User
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}

// DecodeUserTiny decodes the UserTiny instance encoded in resp body.
func (c *Client) DecodeUserTiny(resp *http.Response) (*UserTiny, error) {
	var decoded UserTiny
	err := c.Decoder.Decode(&decoded, resp.Body, resp.Header.Get("Content-Type"))
	return &decoded, err
}
