package routers

import (
	"ginproject/controller/chapter01"
	"ginproject/controller/chapter02"
	"ginproject/controller/chapter03"
	"ginproject/controller/chapter04"
	"ginproject/controller/chapter05"

	"github.com/gin-gonic/gin"
)

func Routers(router *gin.Engine) {
	ch01 := router.Group("/chapter01")
	ch02 := router.Group("/chapter02")
	ch02.Use(chapter05.MiddleWare1) // Router group MiddleWare
	ch03 := router.Group("/chapter03")
	ch04 := router.Group("/chapter04")
	ch05 := router.Group("/chapter05")

	chapter01.Router(ch01)
	chapter02.Router(ch02)
	chapter03.Router(ch03)
	chapter04.Router(ch04)
	chapter05.Router(ch05)
}
