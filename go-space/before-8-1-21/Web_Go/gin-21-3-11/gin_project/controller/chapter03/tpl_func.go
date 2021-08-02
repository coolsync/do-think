package chapter03

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func FuncTpl(ctx *gin.Context) {

	slice_data := []string{"jerry", "paul", "mark"}

	map_d := map[string]interface{}{
		"name": "cherry",
		"age":  20,
	}

	// now_time := time.Now().Format("2006-1-2 15:04:05")
	now_time := time.Now()
	fmt.Println(now_time)

	map_data := map[string]interface{}{
		"name":       "bob",
		"age":        30,
		"slice_data": slice_data,
		"map_d":      map_d,
		"now_time":   now_time,
	}

	ctx.HTML(http.StatusOK, "chapter03/tpl_func.html", map_data)
}

func AddNum(n1, n2 int) int {
	return n1 + n2
}

func SubStr(str string, n int) string {

	str_len := len(str)

	if str_len <= n {
		return str
	}

	if str_len > n {
		n = str_len
	}

	return str[0:n] + "..."
}

func Safe(s string) template.HTML {
	return template.HTML(s)
}
