package domain

import (
	"fmt"
	"mvc/utils"
	"net/http"
)

var (
	users = map[int64]*User{
		123: {Id: 123, First_name: "jo", Last_name: "h", Email: "ssss"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not_found",
	}
}
