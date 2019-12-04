package polo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	polo = "polo"
)

func Polo(c *gin.Context) {
	c.String(http.StatusOK, polo)
}
