package ch04

import "github.com/gin-gonic/gin"

// Ch04
func Routers(ch04_routers *gin.RouterGroup) {
	// Data Binding
	ch04_routers.GET("/to_bind_form", ToBindForm)
	ch04_routers.POST("/do_bind_form", DoBindForm)

	ch04_routers.GET("/to_bind_json", ToBindJson)
	ch04_routers.POST("/do_bind_json", DoBindJson)

	ch04_routers.GET("/bind_query", GetQueryData)
	ch04_routers.GET("/bind_uri/:name/:age/:addr", BindUri)

	// Validator
	ch04_routers.GET("/to_valid", ToValidData)
	ch04_routers.POST("/do_valid", DoValidData)

	// Beego Validator
	ch04_routers.GET("/to_beego_validator", ToBeegoValidator)
	ch04_routers.POST("/do_beego_validator", DoBeegoValidator)
}
