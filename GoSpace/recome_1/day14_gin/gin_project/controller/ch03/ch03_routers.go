package ch03

import "github.com/gin-gonic/gin"

func Routers(ch03_routers *gin.RouterGroup) {
	// Ch03
	// tpl syntax
	ch03_routers.GET("/tpl_syntax1", TplSyntax1)
	ch03_routers.GET("/tpl_syntax2", TplSyntax2)

	// tpl func
	ch03_routers.GET("/tpl_func1", TplFunc1)
	ch03_routers.GET("/tpl_func2", TplFunc2)
	ch03_routers.GET("/consume_tpl_func", ConsumeTplFunc)
}
