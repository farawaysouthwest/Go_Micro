package service

import (
	mocks "go_mvc/mocks"
	"go_mvc/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type testContext struct {
	mockModel *mocks.Model
	service Service
}

func newTestContext(t *testing.T) testContext{
	
	mockModel := &mocks.Model{}

	return testContext{
		mockModel: mockModel,
		service: NewService(mockModel),
	}
}

func TestFetchUser(t *testing.T) {
	
	tc := newTestContext(t)

	tc.mockModel.On("GetUser", mock.Anything).Return(model.User{})

	result := tc.service.FetchUser(123)

	var expected model.User
	assert.IsType(t, expected, result)
}