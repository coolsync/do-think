package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserInfo struct {
// 	Name string `form:"name" json:"name"`
// 	Age  int    `form:"age" json:"age"`
// 	Addr string `form:"addr" json:"addr"`
// }

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/to_json_form.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user UserInfo

	err := ctx.ShouldBind(&user)
	if err != nil {
		fmt.Printf("should bind failed, err: %v\n", err)
		// ctx.JSON(http.StatusOK, gin.H{	// how at ajax show error
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "Failed",
		})
	} else {
		fmt.Println(user)
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "OK",
		})
	}
}
