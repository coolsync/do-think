package ch01

import "github.com/gin-gonic/gin"

func Routers(ch01_routers *gin.RouterGroup) {
	// Ch01
	ch01_routers.GET("/", Hello)
}
