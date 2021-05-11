package main

import (
	"gin2/ch01"
	"gin2/ch02"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create router
	router := gin.Default()

	// Use tmpl Regular, specify 多级目录
	router.LoadHTMLGlob("templates/**/*")

	// Specify static file dir
	router.Static("/static", "static")
	// router.StaticFS("/static", http.Dir("static"))

	// Ch01
	router.GET("/", ch01.Hello)

	// Ch02 Tmpls Render
	router.GET("/user", ch02.User)
	router.GET("/user_info", ch02.UserInfoStruct)
	router.GET("/arr", ch02.ArrayHandler)
	router.GET("/arr_struct", ch02.ArrayAndStruct)
	router.GET("/map", ch02.MapHandler)
	router.GET("/map_struct", ch02.MapAndStruct)

	router.GET("/param1/:name", ch02.Param1)
	router.GET("/param2/*name", ch02.Param2)


	// Listen port
	router.Run(":8090")
}
