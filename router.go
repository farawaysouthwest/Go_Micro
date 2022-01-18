package main

import (
	controller "go_micro/controller"

	"github.com/gin-gonic/gin"
)

type Router struct {
	Router *gin.Engine
	Controller controller.Controller
}


func NewRouter(controller controller.Controller) *Router {

	return &Router{
		Router: gin.Default(),
		Controller: controller,
	}
	
}