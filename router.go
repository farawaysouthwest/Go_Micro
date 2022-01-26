package main

import (
	controller "go_micro/controller"

	"github.com/gin-gonic/gin"
)

type Router interface {

	GetRouter() *gin.Engine

	GetController() controller.Controller

}
type router struct {
	Router *gin.Engine
	Controller controller.Controller
}


func NewRouter(controller controller.Controller, routeEngine *gin.Engine) Router {
	return &router{
		Router: routeEngine,
		Controller: controller,
	}
	
}

func (r router) GetRouter() *gin.Engine {
	return r.Router
}

func (r router) GetController() controller.Controller {
	return r.Controller
}