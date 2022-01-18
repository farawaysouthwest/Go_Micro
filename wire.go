//go:build wireinject
// +build wireinject

package main

import (
	controller "go_micro/controller"
	model "go_micro/model"
	svc "go_micro/service"

	"github.com/google/wire"
)




func InitializeRouter() *Router {
	wire.Build(svc.NewService, model.NewModel, controller.NewController, NewRouter)
	return &Router{}
}