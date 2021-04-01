package handlerpages

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "user.html", nil)
}
