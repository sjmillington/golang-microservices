package domain

import (
	"fmt"
	"net/http"

	"github.com/sjmillington/golang-microservices/mvc/utils"
)

//mock
var (
	users = map[int64]*User{
		123: &User{ID: 1, FistName: "Sam", LastName: "Millington", Email: "someEmail@email.com"},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
