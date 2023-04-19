package users

import (
	"github.com/MELI-jarocha/bookstore_users-api/domain/users"
	"github.com/MELI-jarocha/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestError) {
	userSearch := &users.User{Id: userId}
	if userError := userSearch.Get(); userError != nil {
		return nil, userError
	}
	return userSearch, nil
}
