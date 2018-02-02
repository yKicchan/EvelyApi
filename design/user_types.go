package design

import (
	. "EvelyApi/config"
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var LoginPayload = Type("LoginPayload", func() {
	Description("認証時に受け取るログイン情報")
	Attribute("id", String, "ユーザーID", func() {
		Example("yKicchan")
	})
	Attribute("password", String, "パスワード", func() {
		Example("password")
	})
	Required("id", "password")
})

var EventPayload = Type("EventPayload", func() {
	Description("イベント作成・編集で受け取るイベント情報")
	Attribute("image", String, "イベントのヘッダーイメージ", func() {
		Default("")
		Example("header.jpg")
	})
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
	Attribute("files", ArrayOf(String), "添付資料", func() {
		Example([]string{"doc1.jpg", "doc2.pdf"})
	})
	Attribute("categorys", ArrayOf(String), "カテゴリ", func() {
		MinLength(1)
		MaxLength(len(Categorys))
		Example([]string{C_WORK_CONF, C_FESTIVAL})
	})
	Attribute("mail", String, "連絡先メールアドレス", func() {
		Format("email")
		Example("yKicchanApp@gmail.com")
	})
	Attribute("tel", String, "連絡先電話番号", func() {
		Example("090-1234-5678")
	})
	Attribute("url", String, "URL", func() {
		Format("uri")
		Example("http://comp.ecc.ac.jp/")
	})
	Attribute("schedules", ArrayOf(Schedule), "イベントの開催予定一覧", func() {
		MinLength(1)
	})
	Attribute("noticeRange", Integer, "通知範囲(m)", func() {
		Minimum(100)
		Maximum(MAX_NOTICE_RANGE)
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
	Required("title", "body", "categorys", "schedules", "noticeRange")
})

var Schedule = Type("Schedule", func() {
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

var EmailPayload = Type("EmailPayload", func() {
	Description("新規登録時のメール送信")
	Attribute("email", String, "メールアドレス", func() {
		Format("email")
		Example("yKicchanApp@gmail.com")
	})
	Required("email")
})

var SignupPayload = Type("SignupPayload", func() {
	Description("アカウント作成時に受け取る情報")
	Reference(TokenPayload)
	Attribute("id", String, "ユーザーID", func() {
		Pattern("^[a-zA-Z0-9_]{4,15}$")
		Example("yKicchan")
	})
	Attribute("password", String, "パスワード", func() {
		MinLength(8)
		Example("password")
	})
	Attribute("name", String, "名前", func() {
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
	Attribute("device_token", func() {
		Default("")
	})
	Attribute("instance_id", func() {
		Default("")
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
		Enum("Pending", "OK", "BAN", "Guest")
		Example("OK")
	})
	Required("email", "state")
})

var UserModifyPayload = Type("UserModifyPayload", func() {
	Description("プロフィール編集の時のペイロード")
	Reference(SignupPayload)
	Attribute("name", func() {
		Default("")
	})
	Attribute("icon", String, "アイコン画像", func() {
		Default("")
		Example("icon.png")
	})
	Attribute("email", func() {
		Default("")
	})
	Attribute("tel", func() {
		Default("")
	})
})

var SettingPayload = Type("SettingPayload", func() {
	Description("ユーザー設定を変更する")
	Attribute("preferences", ArrayOf(String), "通知を許可するカテゴリの配列", func() {
		MaxLength(len(Categorys))
		Example([]string{C_WORK_CONF, C_FESTIVAL})
	})
})

var NotifyPayload = Type("NotifyPayload", func() {
	Description("通知のための現在位置情報\nゲストユーザーは通知先となるインスタンスIDを設定する")
	Reference(Location)
	Attribute("lat")
	Attribute("lng")
	Attribute("instance_id", String, "インスタンスID", func() {
		Default("")
		Example("id")
	})
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
	Attribute("files", ArrayOf(String), "レビュー画像など", func() {
		Example([]string{"angle1.jpg", "angle2.jpg"})
	})
	Required("rate", "title", "body")
})

var TokenPayload = Type("TokenPayload", func() {
	Description("通知先を更新・登録するためのデバイストークンとインスタンスIDのペア")
	Attribute("device_token", String, "デバイストークン", func() {
		Example("token")
	})
	Attribute("instance_id", String, "インスタンスID", func() {
		Example("id")
	})
	Required("device_token", "instance_id")
})
