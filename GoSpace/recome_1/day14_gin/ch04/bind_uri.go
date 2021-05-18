package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserInfo struct {
// 	Name string `form:"name" json:"name" uri:"name"`
// 	Age  int    `form:"age" json:"age" uri:"age"`
// 	Addr string `form:"addr" json:"addr" uri:"addr"`
// }

// 统一资源标识符(Uniform Resource Identifier)
func BindUri(ctx *gin.Context) {
	var user UserInfo
	err := ctx.ShouldBindUri(&user)
	if err != nil {
		ctx.String(http.StatusNotFound, "Bind uri failed")
	}

	fmt.Println(user)
	ctx.String(http.StatusOK, "Bind uri OK: %v\n", user)
}
