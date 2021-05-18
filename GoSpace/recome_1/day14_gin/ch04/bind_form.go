package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Name string `form:"name" json:"name" uri:"name"`
	Age  int    `form:"age" json:"age" uri:"age"`
	Addr string `form:"addr" json:"addr" uri:"addr"`
}

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/to_bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	var user_info UserInfo

	err := ctx.ShouldBind(&user_info)
	if err != nil {
		ctx.String(http.StatusNotFound, "Bind Failed")
	}

	fmt.Println(user_info)

	ctx.String(http.StatusOK, "Bind OK!")
}
