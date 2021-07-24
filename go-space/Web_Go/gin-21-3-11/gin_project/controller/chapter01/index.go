package chapter01

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	name := "bob"
	c.HTML(http.StatusOK, "chapter01/index.html", name)

	// res := c.Query("www")
	// fmt.Println("res:", res)
	// c.String(http.StatusOK, res)
}
