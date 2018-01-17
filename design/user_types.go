package design

import (
	. "EvelyApi/config"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var LoginPayload = Type("LoginPayload", func() {
	Description("認証時に受け取るログイン情報")
	Attribute("id", String, "ユーザーID", func() {
		MinLength(1)
		Example("yKicchan")
	})
	Attribute("password", String, "パスワード", func() {
		MinLength(1)
		Example("password")
	})
	Required("id", "password")
})

var EventPayload = Type("EventPayload", func() {
	Description("イベント作成・編集で受け取るイベント情報")
	Attribute("title", String, "イベントの名前", func() {
		MinLength(1)
		MaxLength(30)
		Example("Git勉強会")
	})
	Attribute("body", String, "イベントの詳細", func() {
		MinLength(1)
		MaxLength(1000)
		Example(`初心者でもGitを扱えるようになる勉強会を開催します！
ノートPCを各自持参してください。`)
	})
	Attribute("mail", String, "連絡先メールアドレス", func() {
		Format("email")
		Default("")
		Example("yKicchanApp@gmail.com")
	})
	Attribute("tel", String, "連絡先電話番号", func() {
		Default("")
		Example("090-1234-5678")
	})
	Attribute("url", String, "URL", func() {
		Format("uri")
		Default("")
		Example("http://comp.ecc.ac.jp/")
	})
	Attribute("plans", ArrayOf(Plan), "イベントの開催予定一覧")
	Attribute("noticeRange", Integer, "通知範囲(m)", func() {
		Minimum(100)
		Maximum(MAX_NOTICE_RANGE)
		Default(500)
		Example(500)
	})
	Attribute("scope", String, "公開範囲", func() {
		Enum("public", "private")
		Default("public")
		Example("public")
	})
	Attribute("openFlg", Boolean, "開催中かどうか", func() {
		Default(false)
		Example(false)
	})
	Required("title", "body", "mail", "tel", "url", "plans", "noticeRange", "scope", "openFlg")
})

var Plan = Type("Plan", func() {
	Description("イベントの開催予定情報")
	Attribute("location", Location)
	Attribute("upcomingDate", UpcomingDate)
	Required("location", "upcomingDate")
})

var Location = Type("Location", func() {
	Description("イベントの開催場所")
	Attribute("name", String, "名前", func() {
		MinLength(1)
		MaxLength(50)
		Example("ECCコンピュータ専門学校2303教室")
	})
	Attribute("lat", Number, "緯度", func() {
		Minimum(-90.0)
		Maximum(90.0)
		Example(34.706424)
	})
	Attribute("lng", Number, "経度", func() {
		Minimum(-180.0)
		Maximum(180.0)
		Example(135.50123)
	})
	Required("name", "lat", "lng")
})

var UpcomingDate = Type("UpcomingDate", func() {
	Description("イベントの開催予定日")
	Attribute("startDate", DateTime, "開始日時")
	Attribute("endDate", DateTime, "終了日時")
	Required("startDate", "endDate")
})

var SignupPayload = Type("SignupPayload", func() {
	Description("新規登録時のメール送信")
	Attribute("email", String, "メールアドレス", func() {
		Format("email")
		Example("yKicchanApp@gmail.com")
	})
	Required("email")
})

var UserPayload = Type("UserPayload", func() {
	Description("アカウント作成時に受け取る情報")
	Attribute("id", String, "ユーザーID", func() {
		MinLength(4)
		MaxLength(15)
		Example("yKicchan")
	})
	Attribute("password", String, "パスワード", func() {
		MinLength(8)
		Example("password")
	})
	Attribute("name", String, "名前", func() {
		MinLength(1)
		MaxLength(50)
		Example("きっちゃそ")
	})
	Attribute("mail", String, "メールアドレス", func() {
		Format("email")
		Example("yKicchanApp@gmail.com")
	})
	Attribute("tel", String, "電話番号", func() {
		Default("")
		Example("090-1234-5678")
	})
    Attribute("InstanceID", String, "端末のインスタンスID", func() {
        Default("")
        Example("token")
    })
	Required("id", "password", "name", "mail")
})

var Mail = Type("Mail", func() {
	Description("メールアドレスとその状態")
	Attribute("email", String, "メールアドレス", func() {
		Format("email")
		Example("yKicchanApp@gmail.com")
	})
	Attribute("state", String, "メールアドレスの状態", func() {
		Enum("Pending", "OK", "BAN")
		Example("OK")
	})
	Required("email", "state")
})

var NotifyByInstanceIDPayload = Type("NotifyByInstanceIDPayload", func() {
	Description("インスタンスIDと現在位置情報")
	Reference(Location)
	Attribute("instanceID", String, "通知先となるインスタンスID", func() {
		Example("token")
	})
	Attribute("lat")
	Attribute("lng")
	Required("instanceID", "lat", "lng")
})

var NotifyByUserIDPayload = Type("NotifyByUserIDPayload", func() {
    Description("現在位置情報")
    Reference(Location)
    Attribute("lat")
    Attribute("lng")
    Required("lat", "lng")
})

var PinPayload = Type("PinPayload", func() {
	Description("ピンする/外すイベントのID")
	Attribute("ids", ArrayOf(String), "ピンするイベントのID配列", func() {
        Example([]string{"5a44d5f2775672b659ba00fa", "5a44d5f2775672b659ba00fb"})
	})
	Required("ids")
})

var ReviewPayload = Type("ReviewPayload", func() {
    Description("レビュー投稿で受け取るレビュー情報")
    Attribute("rate", Integer, "レート評価", func() {
        Enum(1, 2, 3, 4, 5)
        Example(4)
    })
    Attribute("title", String, "レビューのタイトル", func() {
		MinLength(1)
		MaxLength(30)
	})
	Attribute("body", String, "レビューの内容", func() {
		MinLength(1)
		MaxLength(5000)
	})
    Required("rate", "title", "body")
})
