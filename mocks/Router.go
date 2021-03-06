// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	controller "go_micro/controller"

	gin "github.com/gin-gonic/gin"

	mock "github.com/stretchr/testify/mock"
)

// Router is an autogenerated mock type for the Router type
type Router struct {
	mock.Mock
}

// GetController provides a mock function with given fields:
func (_m *Router) GetController() controller.Controller {
	ret := _m.Called()

	var r0 controller.Controller
	if rf, ok := ret.Get(0).(func() controller.Controller); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(controller.Controller)
		}
	}

	return r0
}

// GetRouter provides a mock function with given fields:
func (_m *Router) GetRouter() *gin.Engine {
	ret := _m.Called()

	var r0 *gin.Engine
	if rf, ok := ret.Get(0).(func() *gin.Engine); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*gin.Engine)
		}
	}

	return r0
}
