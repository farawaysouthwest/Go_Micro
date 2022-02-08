package controller_test

import (
	"go_micro/controller"
	"go_micro/mocks"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type testContext struct {
	service *mocks.Service
	controller controller.Controller
}

func newTestContext(t *testing.T) (ctx testContext){
	ctx.service = &mocks.Service{}
	ctx.controller = controller.NewController(ctx.service)

	return ctx
}

func Test_Controller_GetUser(t *testing.T) {
	
	tc := newTestContext(t)
	w := httptest.NewRecorder()

	mockGin, _ := gin.CreateTestContext(w)
	
	tc.controller.GetUser(mockGin)

	assert.NotNil(t, w.Result())
}