package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TokenMedia = MediaType("application/vnd.token+json", func() {
	Description("アクセストークン")
	Attributes(func() {
		Attribute("token", String, "アクセストークン", func() {
			Example("Bearer TokenString")
		})
	})
	Required("token")
	View("default", func() {
		Attribute("token")
	})
})

var EventMedia = MediaType("application/vnd.event+json", func() {
	Description("イベント情報")
	Reference(EventPayload)
	Attributes(func() {
		Attribute("id", String, "イベントID", func() {
			Example("20170225-2")
		})
		Attribute("title")
		Attribute("body")
		Attribute("host", UserMedia)
		Attribute("place")
		Attribute("updateDate", DateTime, "最終更新日時")
		Attribute("upcomingDate")
		Attribute("url")
		Attribute("mail")
		Attribute("tel")
	})
	Required("id", "title", "body", "host", "place", "updateDate", "upcomingDate", "url", "mail", "tel")
	View("default", func() {
		Attribute("id")
		Attribute("title")
		Attribute("body")
		Attribute("host", func() {
			View("tiny")
		})
		Attribute("place")
		Attribute("updateDate")
		Attribute("upcomingDate")
		Attribute("url")
		Attribute("mail")
		Attribute("tel")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("title")
		Attribute("host", func() {
			View("tiny")
		})
		Attribute("place")
		Attribute("upcomingDate")
	})
})

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("ユーザー情報")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("mail", Mail)
		Attribute("tel")
	})
	Required("id", "name", "mail", "tel")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("mail")
		Attribute("tel")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("name")
	})
})

var TokenStateMedia = MediaType("application/vnd.token_state+json", func() {
	Description("トークンの状態を返す")
	Attribute("state", String, "状態", func() {
		Enum("Available", "Unavailable")
	})
	Required("state")
	View("default", func() {
		Attribute("state")
	})
})
