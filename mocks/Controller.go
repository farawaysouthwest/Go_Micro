// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// GetUser provides a mock function with given fields: c
func (_m *Controller) GetUser(c *gin.Context) {
	_m.Called(c)
}
