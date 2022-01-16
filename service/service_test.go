package service

import (
	mocks "go_micro/mocks"
	"go_micro/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testContext struct {
	mockModel *mocks.Model
	service Service
}

func newTestContext(t *testing.T) (ctx testContext){
	
	ctx.mockModel = &mocks.Model{}
	ctx.service = NewService(ctx.mockModel)

	return ctx
}

func Test_Service_FetchUser(t *testing.T) {
	
	tc := newTestContext(t)

	tc.mockModel.On("GetUser", mock.Anything).Return(model.User{}, nil)

	result := tc.service.FetchUser(123)

	var expected model.User
	assert.IsType(t, expected, result)
}