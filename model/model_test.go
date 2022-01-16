package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testContext struct {
	model Model
}

func newTestContext(t *testing.T) (ctx testContext){
	
	ctx.model = NewModel()

	return ctx
}

func Test_Model_GetUser(t *testing.T) {

	var id int64 = 123
	var resultUser User

	tc := newTestContext(t)

	result, err := tc.model.GetUser(id)

	assert.Nil(t, err)
	assert.IsType(t, resultUser, result)
}