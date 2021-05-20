package routers

import (
	"gin2/controller/ch01"
	"gin2/controller/ch02"
	"gin2/controller/ch03"
	"gin2/controller/ch04"

	"github.com/gin-gonic/gin"
)

func AllRouters(router *gin.Engine) {
	ch01_router_group := router.Group("/")
	ch02_router_group := router.Group("/ch02")
	ch03_router_group := router.Group("/ch03")
	ch04_router_group := router.Group("/ch04")

	ch01.Routers(ch01_router_group)
	ch02.Routers(ch02_router_group)
	ch03.Routers(ch03_router_group)
	ch04.Routers(ch04_router_group)

}
