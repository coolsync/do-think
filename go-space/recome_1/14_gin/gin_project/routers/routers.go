package routers

import (
	"gin2/controllers/ch01"
	"gin2/controllers/ch02"
	"gin2/controllers/ch03"
	"gin2/controllers/ch04"
	"gin2/controllers/ch05"

	"github.com/gin-gonic/gin"
)

func AllRouters(router *gin.Engine) {
	ch01_router_group := router.Group("/")
	ch02_router_group := router.Group("/ch02")

	ch02_router_group.Use(ch05.MiddleWare1) // Use of routing group middleware

	ch03_router_group := router.Group("/ch03")
	ch04_router_group := router.Group("/ch04")
	ch05_router_group := router.Group("/ch05")

	ch01.Routers(ch01_router_group)
	ch02.Routers(ch02_router_group)
	ch03.Routers(ch03_router_group)
	ch04.Routers(ch04_router_group)
	ch05.Routers(ch05_router_group)
}
