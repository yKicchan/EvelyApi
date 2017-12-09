package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// 認証系、サインアップ、サインインなど
var _ = Resource("auth", func() {
	BasePath("/auth")
	Action("signin", func() {
		Description("サインイン")
		NoSecurity()
		Routing(POST("/signin"))
		Payload(LoginPayload)
		Response(OK, TokenMedia)
		Response(Unauthorized)
	})
})

// イベントに対するアクション
var _ = Resource("events", func() {
	BasePath("/events")

	Action("list", func() {
		Description("イベント複数取得")
		NoSecurity()
		Routing(GET(""))
		Params(func() {
			Param("limit", Integer, "取得件数", func() {
				Minimum(5)
				Maximum(50)
				Default(10)
				Example(30)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(10)
			})
			Param("keyword", String, "キーワード", func() {
				MinLength(1)
				MaxLength(50)
				Example("花澤香菜 Live")
			})
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Required("limit", "offset")
		})
		Response(OK, func() {
			Media(CollectionOf(EventMedia, func() {
				View("default")
				View("tiny")
			}))
		})
		Response(NotFound)
		Response(BadRequest)
	})

	Action("show", func() {
		Description("イベント情報取得")
		NoSecurity()
		Routing(GET("/:user_id/:event_id"))
		Params(func() {
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Param("event_id", String, "イベントID", func() {
				Pattern("^[0-9]{8}-[0-9]+$")
				Example("20170225-2")
			})
		})
		Response(OK, EventMedia)
		Response(NotFound)
	})

	Action("create", func() {
		Description("イベント作成")
		Routing(POST(""))
		Payload(EventPayload)
		Response(Created, EventMedia)
		Response(BadRequest)
	})

	Action("update", func() {
		Description("イベント編集")
		Routing(PUT("/:user_id/:event_id"))
		Params(func() {
			Param("event_id", String, "イベントID", func() {
				Pattern("^[0-9]{8}-[0-9]+$")
				Example("20170225-2")
			})
		})
		Payload(EventPayload)
		Response(OK, EventMedia)
		Response(NotFound)
		Response(BadRequest)
	})

	Action("delete", func() {
		Description("イベント削除")
		Routing(DELETE("/:user_id/:event_id"))
		Params(func() {
			Param("event_id", String, "イベントID", func() {
				Pattern("^[0-9]{8}-[0-9]+$")
				Example("20170225-2")
			})
		})
		Response(OK)
		Response(NotFound)
	})
})

// アカウントに対するアクション
var _ = Resource("users", func() {
	BasePath("/users")
	Action("show", func() {
		Description("アカウント情報取得")
		NoSecurity()
		Routing(GET("/:user_id"))
		Params(func() {
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
		})
		Response(OK, UserMedia)
		Response(BadRequest)
		Response(NotFound)
	})
})

var _ = Resource("swagger", func() {
	NoSecurity()
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Metadata("swagger:generate", "false")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "public/swaggerui")
})
