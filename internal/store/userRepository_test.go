package store_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/varkadov/http-rest-api.git/internal/model"
	"github.com/varkadov/http-rest-api.git/internal/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s, teardown := store.TextStore(t, databaseURL)

	defer teardown("users")

	u, err := s.User().Create(&model.User{
		Email: "user@example.org",
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s, teardown := store.TextStore(t, databaseURL)
	defer teardown("users")

	email := "user@example.org"

	// User not exist
	_, err := s.User().FindByEmail(email)
	assert.Error(t, err)

	_, _ = s.User().Create(&model.User{
		Email: email,
	})

	// User exist
	u, err := s.User().FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, u)
}