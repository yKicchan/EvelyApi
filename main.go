//go:generate goagen bootstrap -d EvelyApi/design

package main

import (
	"EvelyApi/app"
	"EvelyApi/controller"
	. "EvelyApi/middleware"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	mgo "gopkg.in/mgo.v2"
	"log"
)

func main() {
	// Create service
	service := goa.New("EvelyApi")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount security middlewares
	app.UseJWTMiddleware(service, NewJWTMiddleware())

	// DB接続
	session, err := mgo.Dial("mongo")
	if err != nil {
		log.Fatalf("Database initialization failed: %s", err)
	}
	// DB切断
	defer session.Close()
	db := session.DB("sandbox")

	// Mount "actions" controller
	c := controller.NewActionsController(service)
	app.MountActionsController(service, c)
	// Mount "auth" controller
	c2, _ := controller.NewAuthController(service, db)
	app.MountAuthController(service, c2)
	// Mount "swagger" controller
	c3 := controller.NewSwaggerController(service)
	app.MountSwaggerController(service, c3)

	// Start service
	if err := service.ListenAndServe(":8888"); err != nil {
		service.LogError("startup", "err", err)
	}

}
