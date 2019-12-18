package oauth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
	"github.com/sjmillington/golang-microservices/oauth-api/src/api/domain/oauth"
	"github.com/sjmillington/golang-microservices/oauth-api/src/api/services"
)

func CreateAccessToken(c *gin.Context) {

	var request oauth.AccessTokenRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}

	token, err := services.OauthService.CreateAccessToken(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, token)

}

func GetAccessToken(c *gin.Context) {

	accessToken := c.Param("token_id")

	token, err := services.OauthService.GetAccessToken(accessToken)

	if err != nil {
		apiError := errors.NewBadRequestAPIError("invalid token")
		c.JSON(apiError.Status(), apiError)
		return
	}

	c.JSON(http.StatusOK, token)
}
