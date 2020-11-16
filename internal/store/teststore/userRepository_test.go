package teststore

import (
	"github.com/stretchr/testify/assert"
	"github.com/varkadov/http-rest-api.git/internal/model"
	"github.com/varkadov/http-rest-api.git/internal/store"
	"testing"
)

func TestUserRepository_Create(t *testing.T) {
	s := New()
	u := model.TestUser(t)

	assert.NoError(t, s.User().Create(u))
	assert.NotNil(t, u)
}

func TestUserRepository_Find(t *testing.T) {
	s := New()
	email := "user@example.org"

	// Not exist
	_, err := s.User().Find(123)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// Exist
	u := model.TestUser(t)
	u.Email = email
	_ = s.User().Create(u)
	u, err = s.User().Find(u.ID)
	assert.NoError(t, err)
	assert.NotEmpty(t, u)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	s := New()
	email := "user@example.org"

	// Not exist
	_, err := s.User().FindByEmail(email)
	assert.EqualError(t, err, store.ErrRecordNotFound.Error())

	// Exist
	u := model.TestUser(t)
	u.Email = email
	_ = s.User().Create(u)
	u, err = s.User().FindByEmail(u.Email)
	assert.NoError(t, err)
	assert.NotEmpty(t, u)
}
