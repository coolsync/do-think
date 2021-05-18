package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// data verification
// Binding
type Article struct {
	ID          int    `form:"id"`
	Title       string `form:"title" binding:"required"`
	Content     string `form:"content" binding:"min=5"`
	Description string `form:"desc" binding:"max=10"`
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/valid_data.html", nil)
}

func DoValidData(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusNotFound, "get article failed: %v\n", article)
		return
	}
	ctx.String(http.StatusOK, "Ok: %v\n", article)
}
