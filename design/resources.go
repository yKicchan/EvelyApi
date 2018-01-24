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
		Routing(POST("/signin"))
		Payload(LoginPayload)
		Response(OK, TokenMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("signup", func() {
		Description("新規登録")
		Routing(POST("/signup"))
		Payload(UserPayload)
		Response(OK, TokenMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("send_mail", func() {
		Description("新規登録用のメール送信")
		Routing(POST("/signup/send_mail"))
		Payload(EmailPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
	})

	Action("verify_token", func() {
		Description("新規登録時のトークンのチェック")
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
		Routing(GET(""))
		Params(func() {
			Param("limit", Integer, "取得件数", func() {
				Minimum(5)
				Maximum(50)
				Default(10)
				Example(10)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(0)
			})
			Param("keyword", String, "キーワード", func() {
				MaxLength(50)
				Default("")
				Example("花澤香菜 Live")
			})
		})
		Response(OK, CollectionOf(EventMedia))
		Response(NotFound, ErrorMedia)
		Response(BadRequest, ErrorMedia)
	})

	Action("nearby", func() {
		Description("近くのイベントを検索する")
		Routing(GET("/nearby"))
		Params(func() {
			Param("lat", Number, "緯度", func() {
				Minimum(-90.0)
				Maximum(90.0)
				Example(34.706424)
			})
			Param("lng", Number, "経度", func() {
				Minimum(-180.0)
				Maximum(180.0)
				Example(135.50123)
			})
			Param("limit", Integer, "取得件数", func() {
				Minimum(1)
				Maximum(50)
				Default(3)
				Example(3)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(10)
			})
			Param("range", Integer, "検索範囲(半径m)", func() {
				Minimum(10)
				Default(500)
			})
			Required("lat", "lng", "range")
		})
		Response(OK, CollectionOf(EventMedia))
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("イベント情報取得")
		Routing(GET("/detail"))
		Params(func() {
			Param("ids", ArrayOf(String), "詳細を見るイベントのID配列", func() {
				Example([]string{"5a44d5f2775672b659ba00fa", "5a44d5f2775672b659ba00fb"})
			})
			Required("ids")
		})
		Response(OK, CollectionOf(EventMedia))
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("create", func() {
		Description("イベント作成")
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(POST(""))
		Payload(EventPayload)
		Response(Created, EventMedia)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})

	Action("modify", func() {
		Description("イベント編集")
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(PUT("/:event_id"))
		Params(func() {
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
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(DELETE("/:event_id"))
		Params(func() {
			Param("event_id", String, "イベントID", func() {
				Example("5a44d5f2775672b659ba00fa")
			})
		})
		Response(OK)
		Response(Unauthorized, ErrorMedia)
		Response(BadRequest, ErrorMedia)
		Response(Forbidden, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("update", func() {
		Description("イベントの開催フラグを更新する")
		Routing(GET("/update"))
		Response(OK)
	})

	Action("notify", func() {
		Description("近くにイベントがあれば通知する")
		Security(OptionalJWT, func() {
			Scope("api:access")
		})
		Routing(POST("/notify"))
		Payload(NotifyPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("pin", func() {
		Description("ユーザーのピンしたイベント一覧を取得する")
		Routing(GET("/pin/:user_id"))
		Params(func() {
			Param("user_id", String, "ユーザーID", func() {
				Example("yKicchan")
			})
			Param("limit", Integer, "取得件数", func() {
				Minimum(1)
				Maximum(50)
				Default(10)
				Example(10)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(10)
			})
		})
		Response(OK, CollectionOf(EventMedia))
		Response(BadRequest, ErrorMedia)
	})

	Action("my_list", func() {
		Description("自分のイベント一覧を取得する")
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(GET("/my_list"))
		Params(func() {
			Param("limit", Integer, "取得件数", func() {
				Minimum(1)
				Maximum(50)
				Default(10)
				Example(10)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(10)
			})
		})
		Response(OK, CollectionOf(EventMedia))
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})
})

// アカウントに対するアクション
var _ = Resource("users", func() {
	BasePath("/users")

	Action("show", func() {
		Description("アカウント情報取得")
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

	Action("update", func() {
		Description("インスタンスIDの登録・更新\n認証ありで登録ユーザーを、認証なしでゲストユーザーを登録・更新する")
		Security(OptionalJWT, func() {
			Scope("api:access")
		})
		Routing(POST("/update/token"))
		Payload(TokenPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})
})

var _ = Resource("pins", func() {
	BasePath("/pins")
	Security(JWT, func() {
		Scope("api:access")
	})

	Action("on", func() {
		Description("ピンする")
		Routing(PUT("/on"))
		Payload(PinPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})

	Action("off", func() {
		Description("ピンを外す")
		Routing(PUT("/off"))
		Payload(PinPayload)
		Response(OK)
		Response(BadRequest, ErrorMedia)
		Response(Unauthorized, ErrorMedia)
	})
})

var _ = Resource("files", func() {
	BasePath("/files")

	Action("upload", func() {
		Description("ファイルアップロード")
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(POST("/upload"))
		Response(OK, ArrayOf(String))
		Response(BadRequest, ErrorMedia)
	})

	Files("/download/*filename", "public/files/")
})

var _ = Resource("reviews", func() {
	BasePath("/reviews")

	Action("list", func() {
		Description("レビューの一覧取得")
		Routing(GET("/:event_id"))
		Params(func() {
			Param("event_id", String, "イベントID", func() {
				Example("5a44d5f2775672b659ba00fa")
			})
			Param("limit", Integer, "取得件数", func() {
				Minimum(1)
				Maximum(50)
				Default(5)
				Example(5)
			})
			Param("offset", Integer, "除外件数", func() {
				Minimum(0)
				Default(0)
				Example(10)
			})
		})
		Response(OK, CollectionOf(ReviewMedia))
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})

	Action("create", func() {
		Description("レビュー投稿")
		Security(JWT, func() {
			Scope("api:access")
		})
		Routing(POST("/:event_id"))
		Payload(ReviewPayload)
		Params(func() {
			Param("event_id", String, "イベントID", func() {
				Example("5a44d5f2775672b659ba00fa")
			})
		})
		Response(Created, ReviewMedia)
		Response(BadRequest, ErrorMedia)
		Response(NotFound, ErrorMedia)
	})
})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Metadata("swagger:generate", "false")
	Files("/swagger/*filename", "public/swagger/")
	Files("/swaggerui/*filepath", "public/swaggerui/")
})
