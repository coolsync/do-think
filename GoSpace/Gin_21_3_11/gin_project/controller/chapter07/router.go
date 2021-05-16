package chapter07

import "github.com/gin-gonic/gin"

func Router(ch07 *gin.RouterGroup) {
	ch07.GET("/log_data", LogTest)
}
