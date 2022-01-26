//go:build wireinject
// +build wireinject

package main

import (
	controller "go_micro/controller"
	svc "go_micro/service"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)


func InitializeRouter() Router {
	wire.Build(svc.NewService, controller.NewController, svc.NewMockDB, NewRouter, gin.Default)
	return &router{}
}