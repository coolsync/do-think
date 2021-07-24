package chapter03

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Id          int
	Title       string
	Description string
}

func SyntaxTpl(ctx *gin.Context) {
	name := "paul"

	arr := []int{11, 222, 33, 444}

	birth := "2006-01-02"

	article := Article{Id: 1, Title: "符箓传说", Description: " ~符箓传说外传, 不容错过。。。"}

	map_data := map[string]interface{}{
		"name":    name,
		"arr":     arr,
		"birth":   birth,
		"article": article,
	}

	ctx.HTML(
		http.StatusOK,
		"chapter03/template1.html",
		map_data,
	)
}
