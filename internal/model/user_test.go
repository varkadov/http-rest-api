package model_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/varkadov/http-rest-api.git/internal/model"
	"testing"
)

func TestUser_BeforeCreate(t *testing.T) {
	u := model.TestUser(t)

	assert.NoError(t, u.BeforeCreate())
	assert.NotEmpty(t, u.EncryptedPassword)
}
