package ch03

import (
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func TplFunc1(ctx *gin.Context) {
	name := "mark"
	age := 30

	map_data := map[string]interface{}{
		"name": "paul",
		"age":  20,
	}

	sli := []int{1, 2, 3, 4, 5}

	msg := map[string]interface{}{
		"name":     name,
		"age":      age,
		"map_data": map_data,
		"sli":      sli,
	}
	ctx.HTML(http.StatusOK, "ch03/03.tpl_func.html", msg)
}

func TplFunc2(ctx *gin.Context) {

	now_time := time.Now()
	m := map[string]interface{}{
		"now_time": now_time,
	}
	ctx.HTML(http.StatusOK, "ch03/04.tpl_func2.html", m)
}

func ConsumeTplFunc(ctx *gin.Context) {

	ctx.HTML(http.StatusOK, "ch03/05.consume_tpl_func.html", nil)
}

// Custom template function

// router.SetFuncMap(template.FuncMap{
// 	"add_num":   ch03.AddNum,
// 	"str_len":   ch03.SubStr,
// 	"safe_html": ch03.SafeHTML,
// })

// define func
func AddNum(n1, n2 int) int {
	return n1 + n2
}

func SubStr(str string, n int) string {
	str_len := len(str)

	if str_len <= n {
		return str
	}

	return str[:n] + "..."
}

func SafeHTML(s string) template.HTML {
	return template.HTML(s)
}
