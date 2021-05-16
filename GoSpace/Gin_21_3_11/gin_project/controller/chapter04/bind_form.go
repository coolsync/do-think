package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `form:"name" json:"name" uri:"name"`
	Age  int    `form:"age" json:"age" uri:"age"`
	Addr string `form:"addr" json:"addr" uri:"addr"`
}

func ToBindForm(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_form.html", nil)
}

func DoBindForm(ctx *gin.Context) {
	var user User
	err := ctx.ShouldBind(&user)
	fmt.Println(err)
	if err != nil {
		ctx.String(http.StatusNotFound, "bind form failed!")
	}

	fmt.Println(user)

	ctx.String(http.StatusOK, "bind form successful!")
}
