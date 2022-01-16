package main

import (
	controller "go_micro/controller"
	model "go_micro/model"
	svc "go_micro/service"

	"github.com/gin-gonic/gin"
)

func main () {
	// Init router.
	router := gin.Default()
	// Init layers.
	model := model.NewModel()
	service := svc.NewService(model)
	controller := controller.NewController(service)

	// Routes
	router.GET("/user/:userid", controller.GetUser) 

	router.Run(":8080")
}