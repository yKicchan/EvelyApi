package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var TokenMedia = MediaType("application/vnd.token+json", func() {
	Description("アクセストークン")
	Attributes(func() {
		Attribute("token", String, "アクセストークン", func() {
			Example("token string")
		})
	})
	Required("token")
	View("default", func() {
		Attribute("token")
	})
})
