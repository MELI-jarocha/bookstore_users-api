package users

import (
	"github.com/MELI-jarocha/bookstore_users-api/domain/users"
	"github.com/MELI-jarocha/bookstore_users-api/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestError) {
	if error := user.Validate(); error != nil {
		return nil, error
	}

	if error := user.Save(); error != nil {
		return nil, error
	}

	return &user, nil
}
