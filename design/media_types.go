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
			Example("5a44d5f2775672b659ba00fa")
		})
		Attribute("image", String, "ヘッダー画像のURL", func() {
			Example("http://evely.net:8888/download/header.jpg")
		})
		Attribute("title")
		Attribute("body")
		Attribute("files", ArrayOf(String), "添付資料へのURL", func() {
			Example([]string{"http://evely.net:8888/download/doc1.jpg", "http://evely.net:8888/download/doc2.pdf"})
		})
		Attribute("categorys")
		Attribute("host", UserMedia)
		Attribute("mail")
		Attribute("tel")
		Attribute("url")
		Attribute("schedules")
		Attribute("noticeRange")
		Attribute("scope")
		Attribute("openFlg")
		Attribute("isReviewed", Boolean, "レビューの有無")
		Attribute("updateDate", DateTime, "最終更新日時")
		Attribute("createdAt", DateTime, "作成日時")
	})
	Required("id", "image", "title", "body", "files", "categorys", "host", "mail", "tel", "url", "schedules", "noticeRange", "scope", "openFlg", "isReviewed", "updateDate", "createdAt")
	View("default", func() {
		Attribute("id")
		Attribute("image")
		Attribute("title")
		Attribute("body")
		Attribute("files")
		Attribute("categorys")
		Attribute("host", func() {
			View("tiny")
		})
		Attribute("mail")
		Attribute("tel")
		Attribute("url")
		Attribute("schedules")
		Attribute("isReviewed")
		Attribute("updateDate")
		Attribute("createdAt")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("image")
		Attribute("title")
		Attribute("categorys")
		Attribute("host", func() {
			View("tiny")
		})
		Attribute("schedules")
		Attribute("isReviewed")
	})
	View("full", func() {
		Attribute("id")
		Attribute("image")
		Attribute("title")
		Attribute("body")
		Attribute("files")
		Attribute("categorys")
		Attribute("host", func() {
			View("tiny")
		})
		Attribute("mail")
		Attribute("tel")
		Attribute("url")
		Attribute("schedules")
		Attribute("noticeRange")
		Attribute("scope")
		Attribute("openFlg")
		Attribute("isReviewed")
		Attribute("updateDate")
		Attribute("createdAt")
	})
})

var UserMedia = MediaType("application/vnd.user+json", func() {
	Description("ユーザー情報")
	Reference(UserPayload)
	Attributes(func() {
		Attribute("id")
		Attribute("name")
		Attribute("icon", String, "アイコン画像のURL", func() {
			Example("http://evely.net:8888/download/icon.png")
		})
		Attribute("mail", Mail)
		Attribute("tel")
		Attribute("pins", ArrayOf(String), "ピンしているイベントのID配列", func() {
			Example([]string{"5a44d5f2775672b659ba00fa", "5a44d5f2775672b659ba00fb"})
		})
		Attribute("createdAt", DateTime, "作成日時")
	})
	Required("id", "name", "icon", "mail", "tel", "pins", "createdAt")
	View("default", func() {
		Attribute("id")
		Attribute("name")
		Attribute("icon")
		Attribute("mail")
		Attribute("tel")
		Attribute("pins")
		Attribute("createdAt")
	})
	View("tiny", func() {
		Attribute("id")
		Attribute("name")
		Attribute("icon")
	})
})

var EmailMedia = MediaType("application/vnd.email+json", func() {
	Description("メールアドレス")
	Reference(Mail)
	Attribute("email")
	Required("email")
	View("default", func() {
		Attribute("email")
	})
})

var ReviewMedia = MediaType("application/vnd.review+json", func() {
	Description("レビュー")
	Reference(ReviewPayload)
	Attributes(func() {
		Attribute("id", String, "レビューID", func() {
			Example("")
		})
		Attribute("rate")
		Attribute("title")
		Attribute("body")
		Attribute("files", ArrayOf(String), "レビュー画像などのURL", func() {
			Example([]string{"http://evely.net:8888/download/angle1.jpg", "http://evely.net:8888/download/angle2.jpg"})
		})
		Attribute("reviewer", UserMedia)
		Attribute("reviewedAt", DateTime, "レビューした日")
	})
	Required("id", "rate", "title", "body", "files", "reviewer", "reviewedAt")
	View("default", func() {
		Attribute("id")
		Attribute("rate")
		Attribute("title")
		Attribute("body")
		Attribute("files")
		Attribute("reviewer", func() {
			View("tiny")
		})
		Attribute("reviewedAt")
	})
})
