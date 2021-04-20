package chapter11

import "github.com/gin-gonic/gin"

func Router(ch11 *gin.RouterGroup) {
	ch11.GET("/api_axios", ApiAxios)
}
