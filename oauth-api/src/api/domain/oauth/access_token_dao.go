package oauth

import (
	"fmt"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

//mock db
var (
	tokens = make(map[string]*AccessToken, 0)
)

func (at *AccessToken) Save() errors.ApiError {
	at.AccessToken = fmt.Sprintf("USR_%d", at.UserId) //generate: todo.
	tokens[at.AccessToken] = at
	return nil
}

func GetAccessTokenByToken(accessToken string) (*AccessToken, errors.ApiError) {

	token := tokens[accessToken]

	if token == nil {
		return nil, errors.NewNotFoundAPIError("No access token found")
	}

	return token, nil
}
