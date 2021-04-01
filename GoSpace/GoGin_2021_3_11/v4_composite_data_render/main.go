package main

import (
	"ginpro/v4_composite_data_render/handlerpages"

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
	router.GET("/user_info", handlerpages.UserInfoStruct)
	router.GET("/arr", handlerpages.ArrayHandler)
	router.GET("/arrstruct", handlerpages.ArrayAndStructHandler)
	router.GET("/map", handlerpages.MapHandler)
	router.GET("/map_struct", handlerpages.MapAndStructHandler)
	router.GET("/param1/:id", handlerpages.Param1Handler)
	router.GET("/param2/*id", handlerpages.Param2Handler)	// 会把 前面的  / plus

	router.Run(":9000")
}
