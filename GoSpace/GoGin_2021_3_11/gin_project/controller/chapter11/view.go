package chapter11

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApiAxios(ctx *gin.Context) {

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code": 200,
			"msg":  "提交成功",
		})
}
