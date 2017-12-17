//go:generate goagen bootstrap -d EvelyApi/design

package main

import (
	"log"
	"EvelyApi/app"
    "EvelyApi/controller/api"
	. "EvelyApi/middleware"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"labix.org/v2/mgo"
)

func main() {
	// Create service
	service := goa.New("EvelyApi")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middleware
	app.UseJWTMiddleware(service, NewJWTMiddleware())

	// DB接続
	session, err := mgo.Dial("mongo")
	if err != nil {
		log.Fatalf("Database initialization failed: %s", err)
	}
	// DB切断
	defer session.Close()
	db := session.DB("develop")

	// Mount "auth" controller
	c := api.NewAuthController(service, db)
	app.MountAuthController(service, c)
	// Mount "events" controller
	c2 := api.NewEventsController(service, db)
	app.MountEventsController(service, c2)
	// Mount "swagger" controller
	c3 := api.NewSwaggerController(service)
	app.MountSwaggerController(service, c3)
	// Mount "users" controller
	c4 := api.NewUsersController(service, db)
	app.MountUsersController(service, c4)

	// Start service
	if err := service.ListenAndServe(":8888"); err != nil {
		service.LogError("startup", "err", err)
	}

}
