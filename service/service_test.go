package service

import (
	"go_micro/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContext struct {
	mockDB MockDB
	service Service
}

func newTestContext(t *testing.T) (ctx testContext){
	ctx.mockDB = NewMockDB()
	ctx.service = NewService(ctx.mockDB)

	return ctx
}

func Test_Service_GetUser(t *testing.T) {
	
	tc := newTestContext(t)

	result, err := tc.service.GetUser(123)

	var expected model.User
	assert.Nil(t, err)
	assert.IsType(t, expected, result)
}