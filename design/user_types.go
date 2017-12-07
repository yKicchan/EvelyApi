package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// 認証時に受け取るログイン情報
var LoginPayload = Type("LoginPayload", func() {
	Member("id", String, "ユーザーID", func() {
		MinLength(1)
		Example("user")
	})
	Member("password", String, "パスワード", func() {
		MinLength(1)
		Example("123qwe")
	})
	Required("id", "password")
})
