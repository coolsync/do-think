package handlerpages

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {

	name := "hallen"
	c.HTML(200, "index.html", name)
}
