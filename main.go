//go:generate goagen bootstrap -d EvelyApi/design

package main

import (
	"EvelyApi/app"
	. "EvelyApi/config"
	"EvelyApi/controllers/api"
	. "EvelyApi/middleware"
	. "EvelyApi/models"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"labix.org/v2/mgo"
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

	// DB接続
	session, err := mgo.Dial(DB_HOST)
	if err != nil {
		log.Fatalf("Database initialization failed: %s", err)
	}
	// DB切断
	defer session.Close()
	db := NewEvelyDB(session.DB(DB_NAME))

	// Mount security middleware
	jwtm, _ := NewJWTMiddleware(db)
	ojwtm, _ := NewOptionalJWTMiddleware(db)
	app.UseJWTMiddleware(service, jwtm)
	app.UseOptionalJWTMiddleware(service, ojwtm)

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
	// Mount "files" controller
	c5 := api.NewFilesController(service)
	// c5 := api.NewFilesController(service, db)
	app.MountFilesController(service, c5)
	// Mount "pins" controller
	c6 := api.NewPinsController(service, db)
	app.MountPinsController(service, c6)
	// Mount "reviews" controller
	c7 := api.NewReviewsController(service, db)
	app.MountReviewsController(service, c7)

	// Start service
	if err := service.ListenAndServe(":80"); err != nil {
		service.LogError("startup", "err", err)
	}

}
