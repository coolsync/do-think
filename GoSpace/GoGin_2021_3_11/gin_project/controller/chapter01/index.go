package chapter01

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {

	name := "hallen"
	c.HTML(200, "chapter01/index.html", name)
}
