package ch01

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	name := "bob"
	ctx.HTML(http.StatusOK, "index/index.html", name)
}
