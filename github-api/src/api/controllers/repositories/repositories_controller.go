package repositories

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sjmillington/golang-microservices/github-api/src/api/domain/repositories"
	"github.com/sjmillington/golang-microservices/github-api/src/api/services"
	"github.com/sjmillington/golang-microservices/github-api/src/api/utils/errors"
)

func CreateRepo(c *gin.Context) {

	var request repositories.CreateRepoRequest

	//check for valid json
	if err := c.ShouldBindJSON(&request); err != nil {
		apiError := errors.NewBadRequestAPIError("invalid json body")
		c.JSON(apiError.Status(), apiError)
		return
	}

	result, err := services.RepositoryService.CreateRepo(request)

	if err != nil {
		c.JSON(err.Status(), err)
		return
	}

	c.JSON(http.StatusCreated, result)

}
