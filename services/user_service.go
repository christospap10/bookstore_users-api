package services

import (
	"net/http"

	"github.com/christospap10/bookstore_users-api/domain/users"
	"github.com/christospap10/bookstore_users-api/utils/errors"
)

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, &errors.RestErr{
			Status:  http.StatusNotFound,
			Message: "user not found",
		}
	}
	return result, nil
}

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, &errors.RestErr{
			Status:  http.StatusBadRequest,
			Message: "invalid user",
		}
	}
	if err := user.Save(); err != nil {
		return nil, &errors.RestErr{
			Status:  http.StatusInternalServerError,
			Message: "error when trying to save user",
		}

	}
	return &user, nil
}
