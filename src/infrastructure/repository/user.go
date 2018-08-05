package repository

import "github.com/takochuu/go-cleanarchitecture/src/domains"

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create() (*domains.User, error) {
	return &domains.User{}, nil
}
