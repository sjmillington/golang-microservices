package domain

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sjmillington/golang-microservices/mvc-gin/utils"
)

//mock
var (
	users = map[int64]*User{
		123: &User{ID: 1, FistName: "Sam", LastName: "Millington", Email: "someEmail@email.com"},
	}

	UserDao userDaoInterface
)

func init() {
	UserDao = &userDao{}
}

type userDao struct{}

type userDaoInterface interface {
	GetUser(int64) (*User, *utils.ApplicationError)
}

func (u *userDao) GetUser(userId int64) (*User, *utils.ApplicationError) {

	log.Println("We're accessing the database")

	if user := users[userId]; user != nil {
		return user, nil
	}

	return nil, &utils.ApplicationError{
		Message:    fmt.Sprintf("user %v was not found", userId),
		StatusCode: http.StatusNotFound,
		Code:       "not found",
	}
}
