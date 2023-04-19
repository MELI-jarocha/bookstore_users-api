package users

import (
	"github.com/MELI-jarocha/bookstore_users-api/domain/users"
	userServices "github.com/MELI-jarocha/bookstore_users-api/services/users_services"
	"github.com/MELI-jarocha/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateUser(context *gin.Context) {
	var user users.User

	if formatError := context.ShouldBindJSON(&user); formatError != nil {
		restError := errors.MessageBadRequestError("Invalid JSON body")
		context.JSON(restError.Status, restError)
		return
	}

	response, saveError := userServices.CreateUser(user)
	if saveError != nil {
		restError := errors.MessageBadRequestError(saveError.Message)
		context.JSON(restError.Status, restError)
		return
	}

	context.JSON(http.StatusCreated, response)
}
