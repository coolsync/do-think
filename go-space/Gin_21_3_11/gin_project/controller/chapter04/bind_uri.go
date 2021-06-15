package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindUri(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBindUri(&user)
	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound, "Bind uri failed")
	}

	ctx.String(http.StatusOK, "Bind Uri Ok")
}
