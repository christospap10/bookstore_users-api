package users

import (
	"net/http"
	"strconv"

	"github.com/christospap10/bookstore_users-api/domain/users"
	"github.com/christospap10/bookstore_users-api/services"
	"github.com/christospap10/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user users.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.NewBadRequestError("invalid user id!, user id should be a number")
		c.JSON(err.Status, err)
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		err := errors.NewNotFoundError("no user found with given id")
		c.JSON(err.Status, err)
		return
	}

	c.JSON(http.StatusOK, user)
}
