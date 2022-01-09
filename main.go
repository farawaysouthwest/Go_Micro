package main

import (
	controller "go_mvc/controller"
	model "go_mvc/model"
	svc "go_mvc/service"

	"github.com/gin-gonic/gin"
)

func main () {
	// Init router.
	router := gin.Default()
	// Init layers.
	model := model.NewModel()
	service := svc.NewService(model)
	controllers := controller.NewController(service)

	// Routes
	router.GET("/user/:userid", controllers.GetUser) 

	router.Run(":8080")
}