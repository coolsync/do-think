package ch03

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID          int
	Name        string
	Description string
}

func TplSyntax1(ctx *gin.Context) {
	name := "mark"

	arr := []int{1, 2, 3, 4, 5}

	map_data := map[string]interface{}{
		"name":    name,
		"arr":     arr,
		"birth":   "2020-01-02",
	}
	ctx.HTML(http.StatusOK, "ch03/01.tpl_syntax.html", map_data)
}

func TplSyntax2(ctx *gin.Context) {

	article := Article{ID: 1, Name: "Golang Ch03", Description: "about gin course"}

	map_data := map[string]interface{}{
		"article": article,
	}
	ctx.HTML(http.StatusOK, "ch03/02.tpl_syntax.html", map_data)
}
