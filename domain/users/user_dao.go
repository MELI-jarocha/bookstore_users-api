package users

import (
	"fmt"
	"github.com/MELI-jarocha/bookstore_users-api/utils/errors"
)

var (
	usersBD = make(map[int64]*User)
)

func (user *User) Get() *errors.RestError {
	userSearch := usersBD[user.Id]
	if userSearch == nil {
		return errors.MessageNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}

	user.Id = userSearch.Id
	user.FirstName = userSearch.FirstName
	user.LastName = userSearch.LastName
	user.Email = userSearch.Email
	user.DateCreated = userSearch.DateCreated

	return nil
}

func (user *User) Save() *errors.RestError {
	userSearch := usersBD[user.Id]

	if userSearch != nil {
		if userSearch.Email == user.Email {
			return errors.MessageBadRequestError(fmt.Sprintf("Email %s already registered", user.Email))
		}
		return errors.MessageBadRequestError(fmt.Sprintf("User %d already exists", user.Id))
	}

	usersBD[user.Id] = user

	return nil
}
