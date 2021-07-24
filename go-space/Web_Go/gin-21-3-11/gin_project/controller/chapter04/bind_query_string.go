package chapter04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BindQueryString(ctx *gin.Context) {
	var user User

	err := ctx.ShouldBind(&user)
	fmt.Println(user)

	if err != nil {
		ctx.String(http.StatusNotFound, "bind query string failed!")
	}

	ctx.String(http.StatusOK, "bind query string successful!")
}