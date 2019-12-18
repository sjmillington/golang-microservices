package application

import (
	"github.com/sjmillington/golang-microservices/github-api/src/api/controllers/polo"
	"github.com/sjmillington/golang-microservices/oauth-api/src/api/controllers/oauth"
)

func mapUrls() {
	router.GET("/marco", polo.Polo)
	router.POST("/oauth/access_token", oauth.CreateAccessToken)
	router.GET("/oauth/access_token/:token_id", oauth.GetAccessToken)
}
