package chapter03

import "github.com/gin-gonic/gin"

func Router(ch03 *gin.RouterGroup) {
	// chapter03 template statement
	ch03.GET("/tpl_syntax", SyntaxTpl)
	ch03.GET("/tpl_func", FuncTpl)
}
