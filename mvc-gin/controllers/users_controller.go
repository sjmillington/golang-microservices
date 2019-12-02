package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sjmillington/golang-microservices/mvc-gin/utils"

	"github.com/sjmillington/golang-microservices/mvc-gin/services"
)

func GetUser(c *gin.Context) {

	//c.Query will fetch query Params (?caller=231)
	userId, err := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}

		utils.RespondError(c, userErr)
		// Just return the Bad Request to the client
		return
	}

	user, apiErr := services.UsersService.GetUser(userId)

	if apiErr != nil {

		utils.RespondError(c, apiErr)
		// Handle the err and return to the client
		return
	}

	//return user to client
	utils.Respond(c, http.StatusOK, user)
}
