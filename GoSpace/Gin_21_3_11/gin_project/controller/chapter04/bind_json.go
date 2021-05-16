package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ToBindJson(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_json.html", nil)
}

func DoBindJson(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)
	
	fmt.Println(err)
	fmt.Println(user)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"code": 404,
			"msg":  "failed",
		})
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
	})
}
