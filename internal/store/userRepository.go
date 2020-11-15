package store

import "github.com/varkadov/http-rest-api.git/internal/model"

type UserRepository interface {
	Create(*model.User) error
	FindByEmail(string) (*model.User, error)
}
