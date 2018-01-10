package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// 認証系、サインアップ、サインインなど
var _ = Resource("auth", func() {
	BasePath("/auth")

	Action("signin", func() {
		Description("ログイン")
		NoSecurity()
		Routing(POST("/signin"))
		Payload(LoginPayload)
		Response(OK, TokenMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("signup", func() {
		Description("新規登録")
		NoSecurity()
		Routing(POST("/signup"))
		Payload(UserPayload)
		Response(OK, TokenMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("send_mail", func() {
		Description("新規登録用のメール送信")
		NoSecurity()
		Routing(POST("/signup/send_mail"))
		Payload(SignupPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})

	Action("verify_token", func() {
		Description("新規登録時のトークンのチェック")
		NoSecurity()
		Routing(GET("/signup/verify_token"))
		Params(func() {
			Param("token", String, "トークン", func() {
				Example("Token string")
			})
			Required("token")
		})
		Response(OK, EmailMedia)
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})
})

// イベントに対するアクション
var _ = Resource("events", func() {
	BasePath("/events")

	Action("list", func() {
		Description("イベント複数取得")
		NoSecurity()
		Routing(GET(""), GET("/:user_id"))
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
				MaxLength(50)
				Default("")
				Example("花澤香菜 Live")
			})
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Required("limit", "offset")
		})
		Response(OK, CollectionOf(EventMedia))
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
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
				Example("5a44d5f2775672b659ba00fa")
			})
		})
		Response(OK, EventMedia)
        Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("create", func() {
		Description("イベント作成")
		Routing(POST(""))
		Payload(EventPayload)
		Response(Created, EventMedia)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})

	Action("modify", func() {
		Description("イベント編集")
		Routing(PUT("/:user_id/:event_id"))
		Params(func() {
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Param("event_id", String, "イベントID", func() {
				Example("5a44d5f2775672b659ba00fa")
			})
		})
		Payload(EventPayload)
		Response(OK, EventMedia)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
		Response(Forbidden, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("delete", func() {
		Description("イベント削除")
		Routing(DELETE("/:user_id/:event_id"))
		Params(func() {
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Param("event_id", String, "イベントID", func() {
				Example("5a44d5f2775672b659ba00fa")
			})
		})
		Response(OK)
		Response(Unauthorized, ErrorMedia)
		Response(Forbidden, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("update", func() {
		Description("イベントの開催フラグを更新する")
		NoSecurity()
		Routing(GET("/update"))
		Response(OK)
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
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})
})

var _ = Resource("files", func() {
	BasePath("/files")

	Action("upload", func() {
		Description("ファイルアップロード")
		Routing(POST("/upload"))
		Response(OK, CollectionOf(FilePathMedia))
		Response(BadRequest, ErrorMedia)
	})

	Files("/files/*filename", "public/files/")
})

var _ = Resource("swagger", func() {
	NoSecurity()
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Metadata("swagger:generate", "false")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "public/swaggerui/")
})
