//go:generate goagen bootstrap -d github.com/tikasan/goa-code-reading/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/tikasan/goa-code-reading/app"
	"github.com/tikasan/goa-code-reading/controllers"
)

func main() {
	// Create service
	service := goa.New("Sample")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "Sample" controller
	c := controllers.NewSampleController(service)
	app.MountSampleController(service, c)
	// Mount "swagger" controller
	c2 := controllers.NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
