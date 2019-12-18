package oauth

import (
	"strings"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

type AccessTokenRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (r *AccessTokenRequest) Validate() errors.ApiError {
	r.Username = strings.TrimSpace(r.Username)

	if r.Username == "" {
		return errors.NewBadRequestAPIError("invalid username")
	}

	r.Password = strings.TrimSpace(r.Password)

	if r.Password == "" {
		return errors.NewBadRequestAPIError("invalid password")
	}

	return nil
}
