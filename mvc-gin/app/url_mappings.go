package app

import (
	"github.com/sjmillington/golang-microservices/mvc-gin/controllers"
)

func mapUrls() {
	//:user_id is a variable
	router.GET("/users/:user_id", controllers.GetUser)
}
