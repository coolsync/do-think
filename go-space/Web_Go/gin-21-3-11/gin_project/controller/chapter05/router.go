package chapter05

import "github.com/gin-gonic/gin"

func Router(ch05 *gin.RouterGroup) {
	// ch05.GET("/", MiddleWare1)
	// ch05.GET("/", MiddleWare2())
	// ch05.GET("/", MiddleWare3)

	ch05.GET("/auth_Basic", gin.BasicAuth(gin.Accounts{
		"bob":   "123456",
		"paul":  "123",
		"jerry": "1234",
	}), gin.WrapF(WrapFDisc), AuthBasicTest)
}
