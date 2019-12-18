package services

import (
	"time"

	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
	"github.com/sjmillington/golang-microservices/oauth-api/src/api/domain/oauth"
)

//Not in the oauth package. This helps with mocking. No need to mock controllers or providers. Only services + external apis.

type oauthService struct{}

type oauthServiceInterface interface {
	CreateAccessToken(req oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError)
	GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError)
}

var (
	OauthService oauthServiceInterface
)

func init() {
	OauthService = &oauthService{}
}

func (s *oauthService) CreateAccessToken(req oauth.AccessTokenRequest) (*oauth.AccessToken, errors.ApiError) {

	if err := req.Validate(); err != nil {
		return nil, err
	}

	user, err := oauth.GetUserByUsernameAndPassword(req.Username, req.Password)

	if err != nil {
		return nil, err
	}

	token := oauth.AccessToken{
		UserId:  user.Id,
		Expires: time.Now().UTC().Add(1 * time.Hour).Unix(),
	}

	if err := token.Save(); err != nil {
		return nil, err
	}

	return &token, nil
}

func (s *oauthService) GetAccessToken(accessToken string) (*oauth.AccessToken, errors.ApiError) {
	token, err := oauth.GetAccessTokenByToken(accessToken)

	if err != nil {
		return nil, err
	}

	if token.IsExpired() {
		return nil, errors.NewNotFoundAPIError("no access token found")
	}

	return token, err
}
