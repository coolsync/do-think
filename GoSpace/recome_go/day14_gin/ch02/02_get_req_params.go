package ch02

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserInfo struct {
// 	ID   int    `form:"id"`
// 	Name string `form:"name"`
// 	Age  int    `form:"age"`
// 	Addr string `form:"addr"`
// }

// 1 Route with parameters: add parameter values directly to the path
// 使用占位符: ，必须得指定这个路径
// The first case: use a placeholder : , you must specify this path
func Param1(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.String(http.StatusOK, "hello, %s", name)
}

// 使用占位符 *，可以不用匹配这个路径
func Param2(ctx *gin.Context) {
	name := ctx.Param("name")

	ctx.String(http.StatusOK, "hello, %s", name)
}

// 二、带参数的路由：路径中使用 param name
func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name", "mark")

	ctx.String(http.StatusOK, "hello, %s, %s", id, name)
}

func GetQueryArrData(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	ctx.String(http.StatusOK, "hello, %v", ids)
}

func GetQueryMapData(ctx *gin.Context) {
	user := ctx.QueryMap("user")

	ctx.String(http.StatusOK, "hello, %v", user)
}
