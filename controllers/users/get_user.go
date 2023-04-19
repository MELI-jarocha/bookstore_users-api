package users

import (
	userServices "github.com/MELI-jarocha/bookstore_users-api/services/users_services"
	"github.com/MELI-jarocha/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(context *gin.Context) {
	userId, userError := strconv.ParseInt(context.Param("user_id"), 10, 64)
	if userError != nil {
		restError := errors.MessageBadRequestError("User id should be a number")
		context.JSON(restError.Status, restError)
		return
	}

	user, getError := userServices.GetUser(userId)
	if getError != nil {
		context.JSON(getError.Status, getError)
		return
	}

	context.JSON(http.StatusOK, user)
}
