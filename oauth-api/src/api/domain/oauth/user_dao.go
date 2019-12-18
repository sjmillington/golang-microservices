package oauth

import (
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

const (
	getUserByUsernameAndPassword = "SELECT id, username FROM users WHERE username=? AND password=?;"
)

var (
	users = map[string]*User{
		"fede": &User{Username: "fede", Id: 123},
	}
)

func GetUserByUsernameAndPassword(username string, password string) (*User, errors.ApiError) {

	user := users[username]

	if user == nil {
		return nil, errors.NewNotFoundAPIError("no user found with given parameters")
	}

	return user, nil

}
