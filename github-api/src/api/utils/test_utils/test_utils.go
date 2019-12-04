package test_utils

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func GetMockedContext(request *http.Request, response *httptest.ResponseRecorder) *gin.Context {
	c, _ := gin.CreateTestContext(response)
	c.Request = request

	return c
}
