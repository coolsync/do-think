package ch02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserInfo struct {
// 	ID   int    `form:"id"`
// 	Name string `form:"name"`
// 	Age  int    `form:"age"`
// 	Addr string `form:"addr"`
// }

// Parameter Bind
// go to user add page
func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add4.html", nil)
}

// user add
func DoUserAdd4(ctx *gin.Context) {
	var user_info UserInfo

	ctx.ShouldBind(&user_info)
	// ctx.ShouldBindWith()

	fmt.Println(user_info)

	ctx.String(http.StatusOK, "user add ok!")
}
