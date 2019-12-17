package application

import (
	"github.com/sjmillington/golang-microservices/github-api/src/api/controllers/polo"
	"github.com/sjmillington/golang-microservices/github-api/src/api/controllers/repositories"
)

func mapUrls() {
	//for google cloud/AWS
	router.GET("/marco", polo.Polo)
	router.POST("/repository", repositories.CreateRepo)
	router.POST("/repositories", repositories.CreateRepos)
}
