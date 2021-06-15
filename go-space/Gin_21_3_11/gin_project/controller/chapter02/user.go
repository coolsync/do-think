package chapter02

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	c.HTML(http.StatusOK, "chapter02/user.html", nil)
}
