package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetQueryData(ctx *gin.Context) {
	var user_info UserInfo

	err := ctx.ShouldBindQuery(&user_info)
	if err != nil || user_info.Name == "" {
		fmt.Printf("should bind query failed, err: %v\n", err)
		ctx.String(http.StatusOK, "get query data failed")
	} else {
		fmt.Println(user_info, err)
		ctx.String(http.StatusOK, "get query data ok: %v\n", user_info)
	}

}
