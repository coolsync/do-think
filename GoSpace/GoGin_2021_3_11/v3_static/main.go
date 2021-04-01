package main

import (
	"ginpro/v3_static/handlerpages"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// router := gin.New()

	router.LoadHTMLGlob("template/**/*")
	// router.LoadHTMLFiles("index.html, news.html")

	// load static source
	router.Static("/static", "static")

	router.GET("/", handlerpages.Hello)
	router.GET("/user", handlerpages.User)

	router.Run(":9000")
}
