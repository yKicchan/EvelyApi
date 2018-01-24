// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "EvelyApi": Application User Types
//
// Command:
// $ goagen
// --design=EvelyApi/design
// --out=$(GOPATH)/src/EvelyApi
// --version=v1.3.1

package client

import (
	"github.com/goadesign/goa"
	"time"
	"unicode/utf8"
)

// 新規登録時のメール送信
type emailPayload struct {
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
}

// Validate validates the emailPayload type instance.
func (ut *emailPayload) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	return
}

// Publicize creates EmailPayload from emailPayload
func (ut *emailPayload) Publicize() *EmailPayload {
	var pub EmailPayload
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	return &pub
}

// 新規登録時のメール送信
type EmailPayload struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
}

// Validate validates the EmailPayload type instance.
func (ut *EmailPayload) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	return
}

// イベント作成・編集で受け取るイベント情報
type eventPayload struct {
	// イベントの詳細
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	// 添付資料
	Files []string `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// イベントのヘッダーイメージ
	Image *string `form:"image,omitempty" json:"image,omitempty" xml:"image,omitempty"`
	// 連絡先メールアドレス
	Mail *string `form:"mail,omitempty" json:"mail,omitempty" xml:"mail,omitempty"`
	// 通知範囲(m)
	NoticeRange *int `form:"noticeRange,omitempty" json:"noticeRange,omitempty" xml:"noticeRange,omitempty"`
	// 開催中かどうか
	OpenFlg *bool `form:"openFlg,omitempty" json:"openFlg,omitempty" xml:"openFlg,omitempty"`
	// イベントの開催予定一覧
	Schedules []*schedule `form:"schedules,omitempty" json:"schedules,omitempty" xml:"schedules,omitempty"`
	// 公開範囲
	Scope *string `form:"scope,omitempty" json:"scope,omitempty" xml:"scope,omitempty"`
	// 連絡先電話番号
	Tel *string `form:"tel,omitempty" json:"tel,omitempty" xml:"tel,omitempty"`
	// イベントの名前
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
	// URL
	URL *string `form:"url,omitempty" json:"url,omitempty" xml:"url,omitempty"`
}

// Finalize sets the default values for eventPayload type instance.
func (ut *eventPayload) Finalize() {
	var defaultImage = ""
	if ut.Image == nil {
		ut.Image = &defaultImage
	}
	var defaultMail = ""
	if ut.Mail == nil {
		ut.Mail = &defaultMail
	}
	var defaultOpenFlg = false
	if ut.OpenFlg == nil {
		ut.OpenFlg = &defaultOpenFlg
	}
	var defaultScope = "public"
	if ut.Scope == nil {
		ut.Scope = &defaultScope
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
	if ut.Schedules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "schedules"))
	}
	if ut.NoticeRange == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "noticeRange"))
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
	if ut.NoticeRange != nil {
		if *ut.NoticeRange < 100 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.noticeRange`, *ut.NoticeRange, 100, true))
		}
	}
	if ut.NoticeRange != nil {
		if *ut.NoticeRange > 5000 {
			err = goa.MergeErrors(err, goa.InvalidRangeError(`request.noticeRange`, *ut.NoticeRange, 5000, false))
		}
	}
	for _, e := range ut.Schedules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if ut.Scope != nil {
		if !(*ut.Scope == "public" || *ut.Scope == "private") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.scope`, *ut.Scope, []interface{}{"public", "private"}))
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
	if ut.Files != nil {
		pub.Files = ut.Files
	}
	if ut.Image != nil {
		pub.Image = *ut.Image
	}
	if ut.Mail != nil {
		pub.Mail = *ut.Mail
	}
	if ut.NoticeRange != nil {
		pub.NoticeRange = *ut.NoticeRange
	}
	if ut.OpenFlg != nil {
		pub.OpenFlg = *ut.OpenFlg
	}
	if ut.Schedules != nil {
		pub.Schedules = make([]*Schedule, len(ut.Schedules))
		for i2, elem2 := range ut.Schedules {
			pub.Schedules[i2] = elem2.Publicize()
		}
	}
	if ut.Scope != nil {
		pub.Scope = *ut.Scope
	}
	if ut.Tel != nil {
		pub.Tel = *ut.Tel
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
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
	// 添付資料
	Files []string `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// イベントのヘッダーイメージ
	Image string `form:"image" json:"image" xml:"image"`
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
	if ut.Schedules == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "schedules"))
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
	if ut.NoticeRange < 100 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.noticeRange`, ut.NoticeRange, 100, true))
	}
	if ut.NoticeRange > 5000 {
		err = goa.MergeErrors(err, goa.InvalidRangeError(`type.noticeRange`, ut.NoticeRange, 5000, false))
	}
	for _, e := range ut.Schedules {
		if e != nil {
			if err2 := e.Validate(); err2 != nil {
				err = goa.MergeErrors(err, err2)
			}
		}
	}
	if !(ut.Scope == "public" || ut.Scope == "private") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.scope`, ut.Scope, []interface{}{"public", "private"}))
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

// メールアドレスとその状態
type mail struct {
	// メールアドレス
	Email *string `form:"email,omitempty" json:"email,omitempty" xml:"email,omitempty"`
	// メールアドレスの状態
	State *string `form:"state,omitempty" json:"state,omitempty" xml:"state,omitempty"`
}

// Validate validates the mail type instance.
func (ut *mail) Validate() (err error) {
	if ut.Email == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "email"))
	}
	if ut.State == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "state"))
	}
	if ut.Email != nil {
		if err2 := goa.ValidateFormat(goa.FormatEmail, *ut.Email); err2 != nil {
			err = goa.MergeErrors(err, goa.InvalidFormatError(`request.email`, *ut.Email, goa.FormatEmail, err2))
		}
	}
	if ut.State != nil {
		if !(*ut.State == "Pending" || *ut.State == "OK" || *ut.State == "BAN") {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.state`, *ut.State, []interface{}{"Pending", "OK", "BAN"}))
		}
	}
	return
}

// Publicize creates Mail from mail
func (ut *mail) Publicize() *Mail {
	var pub Mail
	if ut.Email != nil {
		pub.Email = *ut.Email
	}
	if ut.State != nil {
		pub.State = *ut.State
	}
	return &pub
}

// メールアドレスとその状態
type Mail struct {
	// メールアドレス
	Email string `form:"email" json:"email" xml:"email"`
	// メールアドレスの状態
	State string `form:"state" json:"state" xml:"state"`
}

// Validate validates the Mail type instance.
func (ut *Mail) Validate() (err error) {
	if ut.Email == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "email"))
	}
	if ut.State == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "state"))
	}
	if err2 := goa.ValidateFormat(goa.FormatEmail, ut.Email); err2 != nil {
		err = goa.MergeErrors(err, goa.InvalidFormatError(`type.email`, ut.Email, goa.FormatEmail, err2))
	}
	if !(ut.State == "Pending" || ut.State == "OK" || ut.State == "BAN") {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.state`, ut.State, []interface{}{"Pending", "OK", "BAN"}))
	}
	return
}

// インスタンスIDと現在位置情報
type notifyByInstanceIDPayload struct {
	// 通知先となるインスタンスID
	InstanceID *string `form:"instanceID,omitempty" json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	// 緯度
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// 経度
	Lng *float64 `form:"lng,omitempty" json:"lng,omitempty" xml:"lng,omitempty"`
}

// Validate validates the notifyByInstanceIDPayload type instance.
func (ut *notifyByInstanceIDPayload) Validate() (err error) {
	if ut.InstanceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "instanceID"))
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
	return
}

// Publicize creates NotifyByInstanceIDPayload from notifyByInstanceIDPayload
func (ut *notifyByInstanceIDPayload) Publicize() *NotifyByInstanceIDPayload {
	var pub NotifyByInstanceIDPayload
	if ut.InstanceID != nil {
		pub.InstanceID = *ut.InstanceID
	}
	if ut.Lat != nil {
		pub.Lat = *ut.Lat
	}
	if ut.Lng != nil {
		pub.Lng = *ut.Lng
	}
	return &pub
}

// インスタンスIDと現在位置情報
type NotifyByInstanceIDPayload struct {
	// 通知先となるインスタンスID
	InstanceID string `form:"instanceID" json:"instanceID" xml:"instanceID"`
	// 緯度
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// 経度
	Lng float64 `form:"lng" json:"lng" xml:"lng"`
}

// Validate validates the NotifyByInstanceIDPayload type instance.
func (ut *NotifyByInstanceIDPayload) Validate() (err error) {
	if ut.InstanceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "instanceID"))
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
	return
}

// 現在位置情報
type notifyByUserIDPayload struct {
	// 緯度
	Lat *float64 `form:"lat,omitempty" json:"lat,omitempty" xml:"lat,omitempty"`
	// 経度
	Lng *float64 `form:"lng,omitempty" json:"lng,omitempty" xml:"lng,omitempty"`
}

// Validate validates the notifyByUserIDPayload type instance.
func (ut *notifyByUserIDPayload) Validate() (err error) {
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
	return
}

// Publicize creates NotifyByUserIDPayload from notifyByUserIDPayload
func (ut *notifyByUserIDPayload) Publicize() *NotifyByUserIDPayload {
	var pub NotifyByUserIDPayload
	if ut.Lat != nil {
		pub.Lat = *ut.Lat
	}
	if ut.Lng != nil {
		pub.Lng = *ut.Lng
	}
	return &pub
}

// 現在位置情報
type NotifyByUserIDPayload struct {
	// 緯度
	Lat float64 `form:"lat" json:"lat" xml:"lat"`
	// 経度
	Lng float64 `form:"lng" json:"lng" xml:"lng"`
}

// Validate validates the NotifyByUserIDPayload type instance.
func (ut *NotifyByUserIDPayload) Validate() (err error) {

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
	return
}

// ピンする/外すイベントのID
type pinPayload struct {
	// ピンするイベントのID配列
	Ids []string `form:"ids,omitempty" json:"ids,omitempty" xml:"ids,omitempty"`
}

// Validate validates the pinPayload type instance.
func (ut *pinPayload) Validate() (err error) {
	if ut.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "ids"))
	}
	return
}

// Publicize creates PinPayload from pinPayload
func (ut *pinPayload) Publicize() *PinPayload {
	var pub PinPayload
	if ut.Ids != nil {
		pub.Ids = ut.Ids
	}
	return &pub
}

// ピンする/外すイベントのID
type PinPayload struct {
	// ピンするイベントのID配列
	Ids []string `form:"ids" json:"ids" xml:"ids"`
}

// Validate validates the PinPayload type instance.
func (ut *PinPayload) Validate() (err error) {
	if ut.Ids == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "ids"))
	}
	return
}

// レビュー投稿で受け取るレビュー情報
type reviewPayload struct {
	// レビューの内容
	Body *string `form:"body,omitempty" json:"body,omitempty" xml:"body,omitempty"`
	// レビュー画像など
	Files []string `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// レート評価
	Rate *int `form:"rate,omitempty" json:"rate,omitempty" xml:"rate,omitempty"`
	// レビューのタイトル
	Title *string `form:"title,omitempty" json:"title,omitempty" xml:"title,omitempty"`
}

// Validate validates the reviewPayload type instance.
func (ut *reviewPayload) Validate() (err error) {
	if ut.Rate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "rate"))
	}
	if ut.Title == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "title"))
	}
	if ut.Body == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "body"))
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) < 1 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 1, true))
		}
	}
	if ut.Body != nil {
		if utf8.RuneCountInString(*ut.Body) > 5000 {
			err = goa.MergeErrors(err, goa.InvalidLengthError(`request.body`, *ut.Body, utf8.RuneCountInString(*ut.Body), 5000, false))
		}
	}
	if ut.Rate != nil {
		if !(*ut.Rate == 1 || *ut.Rate == 2 || *ut.Rate == 3 || *ut.Rate == 4 || *ut.Rate == 5) {
			err = goa.MergeErrors(err, goa.InvalidEnumValueError(`request.rate`, *ut.Rate, []interface{}{1, 2, 3, 4, 5}))
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
	return
}

// Publicize creates ReviewPayload from reviewPayload
func (ut *reviewPayload) Publicize() *ReviewPayload {
	var pub ReviewPayload
	if ut.Body != nil {
		pub.Body = *ut.Body
	}
	if ut.Files != nil {
		pub.Files = ut.Files
	}
	if ut.Rate != nil {
		pub.Rate = *ut.Rate
	}
	if ut.Title != nil {
		pub.Title = *ut.Title
	}
	return &pub
}

// レビュー投稿で受け取るレビュー情報
type ReviewPayload struct {
	// レビューの内容
	Body string `form:"body" json:"body" xml:"body"`
	// レビュー画像など
	Files []string `form:"files,omitempty" json:"files,omitempty" xml:"files,omitempty"`
	// レート評価
	Rate int `form:"rate" json:"rate" xml:"rate"`
	// レビューのタイトル
	Title string `form:"title" json:"title" xml:"title"`
}

// Validate validates the ReviewPayload type instance.
func (ut *ReviewPayload) Validate() (err error) {

	if ut.Title == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "title"))
	}
	if ut.Body == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "body"))
	}
	if utf8.RuneCountInString(ut.Body) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.body`, ut.Body, utf8.RuneCountInString(ut.Body), 1, true))
	}
	if utf8.RuneCountInString(ut.Body) > 5000 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.body`, ut.Body, utf8.RuneCountInString(ut.Body), 5000, false))
	}
	if !(ut.Rate == 1 || ut.Rate == 2 || ut.Rate == 3 || ut.Rate == 4 || ut.Rate == 5) {
		err = goa.MergeErrors(err, goa.InvalidEnumValueError(`type.rate`, ut.Rate, []interface{}{1, 2, 3, 4, 5}))
	}
	if utf8.RuneCountInString(ut.Title) < 1 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.title`, ut.Title, utf8.RuneCountInString(ut.Title), 1, true))
	}
	if utf8.RuneCountInString(ut.Title) > 30 {
		err = goa.MergeErrors(err, goa.InvalidLengthError(`type.title`, ut.Title, utf8.RuneCountInString(ut.Title), 30, false))
	}
	return
}

// イベントの開催予定情報
type schedule struct {
	Location     *location     `form:"location,omitempty" json:"location,omitempty" xml:"location,omitempty"`
	UpcomingDate *upcomingDate `form:"upcomingDate,omitempty" json:"upcomingDate,omitempty" xml:"upcomingDate,omitempty"`
}

// Validate validates the schedule type instance.
func (ut *schedule) Validate() (err error) {
	if ut.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "location"))
	}
	if ut.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "upcomingDate"))
	}
	if ut.Location != nil {
		if err2 := ut.Location.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	if ut.UpcomingDate != nil {
		if err2 := ut.UpcomingDate.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// Publicize creates Schedule from schedule
func (ut *schedule) Publicize() *Schedule {
	var pub Schedule
	if ut.Location != nil {
		pub.Location = ut.Location.Publicize()
	}
	if ut.UpcomingDate != nil {
		pub.UpcomingDate = ut.UpcomingDate.Publicize()
	}
	return &pub
}

// イベントの開催予定情報
type Schedule struct {
	Location     *Location     `form:"location" json:"location" xml:"location"`
	UpcomingDate *UpcomingDate `form:"upcomingDate" json:"upcomingDate" xml:"upcomingDate"`
}

// Validate validates the Schedule type instance.
func (ut *Schedule) Validate() (err error) {
	if ut.Location == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "location"))
	}
	if ut.UpcomingDate == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "upcomingDate"))
	}
	if ut.Location != nil {
		if err2 := ut.Location.Validate(); err2 != nil {
			err = goa.MergeErrors(err, err2)
		}
	}
	return
}

// 通知先を更新・登録するためのデバイストークンとインスタンスIDのペア
type tokenPayload struct {
	// デバイストークン
	DeviceToken *string `form:"device_token,omitempty" json:"device_token,omitempty" xml:"device_token,omitempty"`
	// インスタンスID
	InstanceID *string `form:"instance_id,omitempty" json:"instance_id,omitempty" xml:"instance_id,omitempty"`
}

// Validate validates the tokenPayload type instance.
func (ut *tokenPayload) Validate() (err error) {
	if ut.DeviceToken == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "device_token"))
	}
	if ut.InstanceID == nil {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`request`, "instance_id"))
	}
	return
}

// Publicize creates TokenPayload from tokenPayload
func (ut *tokenPayload) Publicize() *TokenPayload {
	var pub TokenPayload
	if ut.DeviceToken != nil {
		pub.DeviceToken = *ut.DeviceToken
	}
	if ut.InstanceID != nil {
		pub.InstanceID = *ut.InstanceID
	}
	return &pub
}

// 通知先を更新・登録するためのデバイストークンとインスタンスIDのペア
type TokenPayload struct {
	// デバイストークン
	DeviceToken string `form:"device_token" json:"device_token" xml:"device_token"`
	// インスタンスID
	InstanceID string `form:"instance_id" json:"instance_id" xml:"instance_id"`
}

// Validate validates the TokenPayload type instance.
func (ut *TokenPayload) Validate() (err error) {
	if ut.DeviceToken == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "device_token"))
	}
	if ut.InstanceID == "" {
		err = goa.MergeErrors(err, goa.MissingAttributeError(`type`, "instance_id"))
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
	// デバイストークン
	DeviceToken *string `form:"device_token,omitempty" json:"device_token,omitempty" xml:"device_token,omitempty"`
	// アイコン画像
	Icon *string `form:"icon,omitempty" json:"icon,omitempty" xml:"icon,omitempty"`
	// ユーザーID
	ID *string `form:"id,omitempty" json:"id,omitempty" xml:"id,omitempty"`
	// インスタンスID
	InstanceID *string `form:"instance_id,omitempty" json:"instance_id,omitempty" xml:"instance_id,omitempty"`
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
	var defaultIcon = ""
	if ut.Icon == nil {
		ut.Icon = &defaultIcon
	}
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
	if ut.DeviceToken != nil {
		pub.DeviceToken = ut.DeviceToken
	}
	if ut.Icon != nil {
		pub.Icon = *ut.Icon
	}
	if ut.ID != nil {
		pub.ID = *ut.ID
	}
	if ut.InstanceID != nil {
		pub.InstanceID = ut.InstanceID
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
	// デバイストークン
	DeviceToken *string `form:"device_token,omitempty" json:"device_token,omitempty" xml:"device_token,omitempty"`
	// アイコン画像
	Icon string `form:"icon" json:"icon" xml:"icon"`
	// ユーザーID
	ID string `form:"id" json:"id" xml:"id"`
	// インスタンスID
	InstanceID *string `form:"instance_id,omitempty" json:"instance_id,omitempty" xml:"instance_id,omitempty"`
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
