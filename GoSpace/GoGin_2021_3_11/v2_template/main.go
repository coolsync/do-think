package main

import "github.com/gin-gonic/gin"

func Hello(c *gin.Context) {

	c.HTML(200, "index.html", nil)

}

func main() {
	router := gin.Default()

	// router := gin.New()

	// router.GET("/", func(c *gin.Context) {
	// 	c.String(200, "hello gin")
	// })

	router.GET("/", Hello)

	router.LoadHTMLGlob("template/*")
	// router.LoadHTMLFiles("index.html, news.html")

	router.Run(":9000")
}
