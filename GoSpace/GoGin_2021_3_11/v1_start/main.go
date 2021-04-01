package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {
	c.String(200, "hello gin")
}

func main() {
	router := gin.Default()

	// router := gin.New()

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "hello gin")
	// })

	router.GET("/user", Hello)

	router.Run(":9000")
}
