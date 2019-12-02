package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/sjmillington/golang-microservices/mvc/utils"

	"github.com/sjmillington/golang-microservices/mvc/services"
)

func GetUser(resp http.ResponseWriter, req *http.Request) {

	userId, err := (strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64))
	if err != nil {
		userErr := &utils.ApplicationError{
			Message:    "user_id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad request",
		}
		resp.WriteHeader(userErr.StatusCode)
		jsonValue, _ := json.Marshal(userErr)
		resp.Write(jsonValue)
		// Just return the Bad Request to the client
		return
	}

	user, apErr := services.UsersService.GetUser(userId)

	if apErr != nil {

		resp.WriteHeader(apErr.StatusCode)
		jsonValue, _ := json.Marshal(apErr)
		resp.Write(jsonValue)
		// Handle the err and return to the client
		return
	}

	//return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
