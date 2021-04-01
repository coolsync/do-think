package chapter04

import "github.com/gin-gonic/gin"

// Chapter04
func Router(ch04 *gin.RouterGroup) {
	// Bind Form
	ch04.GET("/to_bind_form", ToBindForm)
	ch04.POST("/do_bind_form", DoBindForm)

	// Bind query string
	ch04.GET("/bind_query_string", BindQueryString)

	// Bind Json
	ch04.GET("/to_bind_json", ToBindJson)
	ch04.POST("/do_bind_json", DoBindJson)

	// Bind Uri
	ch04.GET("/bind_uri/:name/:age/:addr", BindUri)

	// Bind Valid data
	ch04.GET("/to_valid", ToBindValid)
	ch04.POST("/do_valid", DoBindValid)
}
