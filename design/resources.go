package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("actions", func() {
	BasePath("/actions")
	Security(JWT, func() {
		Scope("api:access")
	})
	Action("ping", func() {
		Description("導通確認")
		Routing(GET("/ping"))
		Response(OK)
		Response(Unauthorized)
	})
})

var _ = Resource("auth", func() {
	BasePath("/auth")

	Security(JWT, func() {
		Scope("api:access")
	})

	Action("signin", func() {
		Description("サインイン")
		NoSecurity()
		Routing(POST("/signin"))
		Payload(LoginPayload)
		Response(OK, TokenMedia)
		Response(Unauthorized)
	})
})

// swagger
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET") // Allow all origins to retrieve the Swagger JSON (CORS)
	})
	Metadata("swagger:generate", "false")
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swaggerui/*filepath", "public/swaggerui")
})
