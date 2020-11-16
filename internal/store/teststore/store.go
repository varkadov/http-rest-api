package teststore

import (
	"github.com/varkadov/http-rest-api.git/internal/model"
	"github.com/varkadov/http-rest-api.git/internal/store"
)

type Store struct {
	userRepository *UserRepository
}

func New() *Store {
	return &Store{}
}

func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
		users: make(map[int]*model.User),
	}
	return s.User()
}
