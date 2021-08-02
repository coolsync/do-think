package ch05

import "github.com/gin-gonic/gin"

func Routers(ch05_routers *gin.RouterGroup) {
	// Use of BasicAuth middleware
	ch05_routers.GET("/auth", gin.BasicAuth(gin.Accounts{
		"bob":   "123456",
		"paul":  "123",
		"jerry": "1234",
	}), gin.WrapF(WrapFHandler), AuthBasicHandler)
}
