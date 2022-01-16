package controller

import (
	"go_micro/mocks"
	"go_micro/model"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testContext struct {
	service *mocks.Service
	controller Controller
}

func newTestContext(t *testing.T) (ctx testContext){
	
	ctx.service = &mocks.Service{}
	ctx.controller = NewController(ctx.service)

	return ctx
}

func Test_Controller_GetUser(t *testing.T) {

	tc := newTestContext(t)
	w := httptest.NewRecorder()
	gin, _ := gin.CreateTestContext(w)

	tc.service.On("FetchUser", mock.Anything).Return(model.User{})

	tc.controller.GetUser(gin)

	assert.NotNil(t, w.Result())
}