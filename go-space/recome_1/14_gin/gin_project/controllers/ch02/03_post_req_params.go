package ch02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// type UserInfo struct {
// 	ID   int    `form:"id"`
// 	Name string `form:"name"`
// 	Age  int    `form:"age"`
// 	Addr string `form:"addr"`
// }

// go to user add page
func ToUserAdd(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add.html", nil)
}

// user add, post
func DoUserAdd(ctx *gin.Context) {
	username := ctx.PostForm("username")
	passwd := ctx.PostForm("passwd")

	fmt.Println(username, passwd)

	ctx.String(http.StatusOK, "user add ok!")
}

// go to user add page
func ToUserAdd2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add2.html", nil)
}

// user add, post
func DoUserAdd2(ctx *gin.Context) {
	username := ctx.PostForm("username")
	passwd := ctx.PostForm("passwd")

	age := ctx.DefaultPostForm("age", "18")
	loves := ctx.PostFormArray("love")
	user := ctx.PostFormMap("user")

	fmt.Println(username)
	fmt.Println(passwd)
	fmt.Println(age)
	fmt.Println(loves)
	fmt.Println(user) // map[addr: ... , phone: ... ]

	ctx.String(http.StatusOK, "user add ok!")
}

// Ajax interactive
// go to ajax add user page
func ToUserAdd3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add3.html", nil)
}

// user add, ajax post
func DoUserAdd3(ctx *gin.Context) {
	username := ctx.PostForm("username")
	passwd := ctx.PostForm("passwd")

	fmt.Println(username)
	fmt.Println(passwd)

	map_data := map[string]interface{}{
		"code": 400,
		"msg":  "FAIL",
	}
	if username == "" || passwd == "" {
		ctx.JSON(http.StatusOK, map_data)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "OK",
		})
	}
}
