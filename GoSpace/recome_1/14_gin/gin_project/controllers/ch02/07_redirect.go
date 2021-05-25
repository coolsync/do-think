package ch02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RedirectA(ctx *gin.Context) {
	fmt.Println("Router A")
	ctx.Redirect(http.StatusFound, "/redirect_b") // 302
	// ctx.Redirect(http.StatusFound, "https://cn.bing.com/") // 302

}

func RedirectB(ctx *gin.Context) {
	fmt.Println("Router B")
	ctx.String(http.StatusOK, "Router B")
}
