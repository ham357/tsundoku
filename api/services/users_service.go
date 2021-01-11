package services

import (
	"github.com/ham357/tsundoku/api/domain/users"
	"github.com/ham357/tsundoku/api/utils/errors"
)

// CreateUser - Service
func CreateUser(user users.User) (*users.User, *errors.ApiErr) {
	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}
