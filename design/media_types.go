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
        Attribute("mail")
        Attribute("tel")
        Attribute("url")
        Attribute("plans")
        Attribute("noticeRange")
        Attribute("scope")
        Attribute("openFlg")
        Attribute("updateDate", DateTime, "最終更新日時")
        Attribute("createdAt")
    })
    Required("id", "title", "body", "host", "mail", "tel", "url", "plans", "noticeRange", "scope", "openFlg", "updateDate", "createdAt")
    View("default", func() {
        Attribute("id")
        Attribute("title")
        Attribute("body")
        Attribute("host", func() {
            View("tiny")
        })
        Attribute("mail")
        Attribute("tel")
        Attribute("url")
        Attribute("plans")
        Attribute("noticeRange")
        Attribute("scope")
        Attribute("openFlg")
        Attribute("updateDate")
        Attribute("createdAt")
    })
    View("tiny", func() {
        Attribute("id")
        Attribute("title")
        Attribute("host", func() {
            View("tiny")
        })
        Attribute("plans")
        Attribute("noticeRange")
        Attribute("scope")
        Attribute("openFlg")
        Attribute("updateDate", DateTime, "最終更新日時")
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

var EmailMedia = MediaType("application/vnd.email+json", func() {
    Description("メールアドレス")
    Reference(Mail)
    Attribute("email")
    Required("email")
    View("default", func() {
        Attribute("email")
    })
})
