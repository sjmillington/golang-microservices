package application

import (
	"github.com/gin-gonic/gin"
	"github.com/sjmillington/golang-microservices/github-api/src/api/log"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartApp() {

	log.Info("about to start the url mappings", "step:1", "status:pending", "wrong")

	mapUrls()

	log.Info("URLs successfully mapped")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
