package chapter01

import "github.com/gin-gonic/gin"

func Router(ch01 *gin.RouterGroup) {
	ch01.GET("/", Hello)
} 