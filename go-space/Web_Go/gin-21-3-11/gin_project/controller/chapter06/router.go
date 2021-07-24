package chapter06

import "github.com/gin-gonic/gin"

func Router(ch06 *gin.RouterGroup) {
	ch06.GET("/gorm_data", GormTest)
}
