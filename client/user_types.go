// Code generated by goagen v1.3.0, DO NOT EDIT.
//
// API "EvelyApi": Application User Types
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.0

package client

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// イベント作成・編集で受け取るイベント情報
type eventPayload struct {
	// イベントの詳細
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	// 連絡先メールアドレス
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// 開催場所
	Place *location `form:"place,omitempty" json:"place,omitempty" xml:"place,omitempty"`
	// 連絡先電話番号
	Tel *string `form:"tel,omitempty" json:"tel,omitempty" xml:"tel,omitempty"`
	// イベントの名前
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// 開催予定日
	UpcomingDate *upcomingDate `form:"upcomingDate,omitempty" json:"upcomingDate,omitempty" xml:"upcomingDate,omitempty"`
	// URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Finalize sets the default values for eventPayload type instance.
func (ut *eventPayload) Finalize() {
	var defaultMail = ""
	if ut.Mail == nil {
		ut.Mail = &defaultMail
	}
	var defaultTel = ""
	if ut.Tel == nil {
		ut.Tel = &defaultTel
	}
	var defaultURL = ""
	if ut.URL == nil {
		ut.URL = &defaultURL
	}
}

// Validate validates the eventPayload type instance.
func (ut *eventPayload) Validate() (err error) {
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "body"))
	}
	if ut.Place == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "place"))
	}
	if ut.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "upcomingDate"))
	}
	if ut.URL == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "url"))
	}
	if ut.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "mail"))
	}
	if ut.Tel == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "tel"))
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 1, true))
		}
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) > 1000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 1000, false))
		}
	}
	if ut.Mail != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Mail); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.mail`, *ut.Mail, goa.FormatEmail, err2))
		}
	}
	if ut.Place != nil {
		if err2 := ut.Place.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 1, true))
		}
	}
	if ut.Title != nil {
		if utf8.RuneCountInString(*ut.Title) > 30 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.title`, *ut.Title, utf8.RuneCountInString(*ut.Title), 30, false))
		}
	}
	if ut.UpcomingDate != nil {
		if err2 := ut.UpcomingDate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.URL != nil {
		if err2 := goa.ValidateFormat(goa.FormatURI, *ut.URL); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.url`, *ut.URL, goa.FormatURI, err2))
		}
	}
	return
}

// Publicize creates EventPayload from eventPayload
func (ut *eventPayload) Publicize() *EventPayload {
	var pub EventPayload
	if ut.Body != nil {
		pub.Body = *ut.Body
	}
	if ut.Mail != nil {
		pub.Mail = *ut.Mail
	}
	if ut.Place != nil {
		pub.Place = ut.Place.Publicize()
	}
	if ut.Tel != nil {
		pub.Tel = *ut.Tel
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	if ut.UpcomingDate != nil {
		pub.UpcomingDate = ut.UpcomingDate.Publicize()
	}
	if ut.URL != nil {
		pub.URL = *ut.URL
	}
	return &pub
}

// イベント作成・編集で受け取るイベント情報
type EventPayload struct {
	// イベントの詳細
	Body string `form:"body" json:"body" xml:"body"`
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
	// URL
	URL string `form:"url" json:"url" xml:"url"`
}

// Validate validates the EventPayload type instance.
func (ut *EventPayload) Validate() (err error) {
	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "body"))
	}
	if ut.Place == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "place"))
	}
	if ut.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "upcomingDate"))
	}
	if ut.URL == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "url"))
	}
	if ut.Mail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "mail"))
	}
	if ut.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "tel"))
	}
	if utf8.RuneCountInString(ut.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.body`, ut.Body, utf8.RuneCountInString(ut.Body), 1, true))
	}
	if utf8.RuneCountInString(ut.Body) > 1000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.body`, ut.Body, utf8.RuneCountInString(ut.Body), 1000, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Mail); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.mail`, ut.Mail, goa.FormatEmail, err2))
	}
	if ut.Place != nil {
		if err2 := ut.Place.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if utf8.RuneCountInString(ut.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.title`, ut.Title, utf8.RuneCountInString(ut.Title), 1, true))
	}
	if utf8.RuneCountInString(ut.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.title`, ut.Title, utf8.RuneCountInString(ut.Title), 30, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatURI, ut.URL); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.url`, ut.URL, goa.FormatURI, err2))
	}
	return
}

// イベントの開催場所
type location struct {
	// 緯度
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// 経度
	Lng *float64 `form:"lng,omitempty" json:"lng,omitempty" xml:"lng,omitempty"`
	// 名前
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
}

// Validate validates the location type instance.
func (ut *location) Validate() (err error) {
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Lat == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "lat"))
	}
	if ut.Lng == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "lng"))
	}
	if ut.Lat != nil {
		if *ut.Lat < -90.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.lat`, *ut.Lat, -90.000000, true))
		}
	}
	if ut.Lat != nil {
		if *ut.Lat > 90.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.lat`, *ut.Lat, 90.000000, false))
		}
	}
	if ut.Lng != nil {
		if *ut.Lng < -180.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.lng`, *ut.Lng, -180.000000, true))
		}
	}
	if ut.Lng != nil {
		if *ut.Lng > 180.000000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.lng`, *ut.Lng, 180.000000, false))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 1, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 50, false))
		}
	}
	return
}

// Publicize creates Location from location
func (ut *location) Publicize() *Location {
	var pub Location
	if ut.Lat != nil {
		pub.Lat = *ut.Lat
	}
	if ut.Lng != nil {
		pub.Lng = *ut.Lng
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	return &pub
}

// イベントの開催場所
type Location struct {
	// 緯度
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// 経度
	Lng float64 `form:"lng" json:"lng" xml:"lng"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
}

// Validate validates the Location type instance.
func (ut *Location) Validate() (err error) {
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}

	if ut.Lat < -90.000000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.lat`, ut.Lat, -90.000000, true))
	}
	if ut.Lat > 90.000000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.lat`, ut.Lat, 90.000000, false))
	}
	if ut.Lng < -180.000000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.lng`, ut.Lng, -180.000000, true))
	}
	if ut.Lng > 180.000000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.lng`, ut.Lng, 180.000000, false))
	}
	if utf8.RuneCountInString(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 1, true))
	}
	if utf8.RuneCountInString(ut.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 50, false))
	}
	return
}

// 認証時に受け取るログイン情報
type loginPayload struct {
	// ユーザーID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
}

// Validate validates the loginPayload type instance.
func (ut *loginPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.ID != nil {
		if utf8.RuneCountInString(*ut.ID) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.id`, *ut.ID, utf8.RuneCountInString(*ut.ID), 1, true))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 1, true))
		}
	}
	return
}

// Publicize creates LoginPayload from loginPayload
func (ut *loginPayload) Publicize() *LoginPayload {
	var pub LoginPayload
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	return &pub
}

// 認証時に受け取るログイン情報
type LoginPayload struct {
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// パスワード
	Password string `form:"password" json:"password" xml:"password"`
}

// Validate validates the LoginPayload type instance.
func (ut *LoginPayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "id"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if utf8.RuneCountInString(ut.ID) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.id`, ut.ID, utf8.RuneCountInString(ut.ID), 1, true))
	}
	if utf8.RuneCountInString(ut.Password) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 1, true))
	}
	return
}

// イベントの開催予定日
type upcomingDate struct {
	// 終了日時
	EndDate *time.Time `form:"endDate,omitempty" json:"endDate,omitempty" xml:"endDate,omitempty"`
	// 開始日時
	StartDate *time.Time `form:"startDate,omitempty" json:"startDate,omitempty" xml:"startDate,omitempty"`
}

// Validate validates the upcomingDate type instance.
func (ut *upcomingDate) Validate() (err error) {
	if ut.StartDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "startDate"))
	}
	if ut.EndDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "endDate"))
	}
	return
}

// Publicize creates UpcomingDate from upcomingDate
func (ut *upcomingDate) Publicize() *UpcomingDate {
	var pub UpcomingDate
	if ut.EndDate != nil {
		pub.EndDate = *ut.EndDate
	}
	if ut.StartDate != nil {
		pub.StartDate = *ut.StartDate
	}
	return &pub
}

// イベントの開催予定日
type UpcomingDate struct {
	// 終了日時
	EndDate time.Time `form:"endDate" json:"endDate" xml:"endDate"`
	// 開始日時
	StartDate time.Time `form:"startDate" json:"startDate" xml:"startDate"`
}

// Validate validates the UpcomingDate type instance.
func (ut *UpcomingDate) Validate() (err error) {

	return
}

// アカウント作成時に受け取る情報
type userPayload struct {
	// ユーザーID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// メールアドレス
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// 名前
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// パスワード
	Password *string `form:"password,omitempty" json:"password,omitempty" xml:"password,omitempty"`
	// 電話番号
	Tel *string `form:"tel,omitempty" json:"tel,omitempty" xml:"tel,omitempty"`
}

// Finalize sets the default values for userPayload type instance.
func (ut *userPayload) Finalize() {
	var defaultTel = ""
	if ut.Tel == nil {
		ut.Tel = &defaultTel
	}
}

// Validate validates the userPayload type instance.
func (ut *userPayload) Validate() (err error) {
	if ut.ID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "id"))
	}
	if ut.Password == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "password"))
	}
	if ut.Name == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "name"))
	}
	if ut.Mail == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "mail"))
	}
	if ut.Tel == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "tel"))
	}
	if ut.ID != nil {
		if utf8.RuneCountInString(*ut.ID) < 4 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.id`, *ut.ID, utf8.RuneCountInString(*ut.ID), 4, true))
		}
	}
	if ut.ID != nil {
		if utf8.RuneCountInString(*ut.ID) > 15 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.id`, *ut.ID, utf8.RuneCountInString(*ut.ID), 15, false))
		}
	}
	if ut.Mail != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Mail); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.mail`, *ut.Mail, goa.FormatEmail, err2))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 1, true))
		}
	}
	if ut.Name != nil {
		if utf8.RuneCountInString(*ut.Name) > 50 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.name`, *ut.Name, utf8.RuneCountInString(*ut.Name), 50, false))
		}
	}
	if ut.Password != nil {
		if utf8.RuneCountInString(*ut.Password) < 8 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.password`, *ut.Password, utf8.RuneCountInString(*ut.Password), 8, true))
		}
	}
	return
}

// Publicize creates UserPayload from userPayload
func (ut *userPayload) Publicize() *UserPayload {
	var pub UserPayload
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.Mail != nil {
		pub.Mail = *ut.Mail
	}
	if ut.Name != nil {
		pub.Name = *ut.Name
	}
	if ut.Password != nil {
		pub.Password = *ut.Password
	}
	if ut.Tel != nil {
		pub.Tel = *ut.Tel
	}
	return &pub
}

// アカウント作成時に受け取る情報
type UserPayload struct {
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// メールアドレス
	Mail string `form:"mail" json:"mail" xml:"mail"`
	// 名前
	Name string `form:"name" json:"name" xml:"name"`
	// パスワード
	Password string `form:"password" json:"password" xml:"password"`
	// 電話番号
	Tel string `form:"tel" json:"tel" xml:"tel"`
}

// Validate validates the UserPayload type instance.
func (ut *UserPayload) Validate() (err error) {
	if ut.ID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "id"))
	}
	if ut.Password == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "password"))
	}
	if ut.Name == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "name"))
	}
	if ut.Mail == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "mail"))
	}
	if ut.Tel == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "tel"))
	}
	if utf8.RuneCountInString(ut.ID) < 4 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.id`, ut.ID, utf8.RuneCountInString(ut.ID), 4, true))
	}
	if utf8.RuneCountInString(ut.ID) > 15 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.id`, ut.ID, utf8.RuneCountInString(ut.ID), 15, false))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Mail); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.mail`, ut.Mail, goa.FormatEmail, err2))
	}
	if utf8.RuneCountInString(ut.Name) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 1, true))
	}
	if utf8.RuneCountInString(ut.Name) > 50 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.name`, ut.Name, utf8.RuneCountInString(ut.Name), 50, false))
	}
	if utf8.RuneCountInString(ut.Password) < 8 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.password`, ut.Password, utf8.RuneCountInString(ut.Password), 8, true))
	}
	return
}
