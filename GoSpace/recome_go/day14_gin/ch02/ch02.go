package ch02

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID   int    `form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
	Addr string `form:"addr"`
}

func User(ctx *gin.Context) {
	name := "mark"
	ctx.HTML(http.StatusOK, "user/user.html", name)
}

// Struct Render
func UserInfoStruct(ctx *gin.Context) {
	var user_info UserInfo
	user_info.ID = 1
	user_info.Name = "paul"
	user_info.Age = 18
	user_info.Addr = "xxx"

	user_info2 := UserInfo{ID: 2, Name: "jerry", Age: 19, Addr: "xxxx"}

	ctx.HTML(http.StatusOK, "ch02/user_info.html", user_info2)
}

// Array Render
func ArrayHandler(ctx *gin.Context) {
	arr := [3]int{1, 2, 3}

	ctx.HTML(http.StatusOK, "ch02/arr.html", arr)
}

// array and sruct render
func ArrayAndStruct(ctx *gin.Context) {
	arr_struct := []UserInfo{
		{ID: 3, Name: "tom", Age: 19, Addr: "xxxx"},
		{ID: 4, Name: "bob", Age: 19, Addr: "xxxx"},
	}

	ctx.HTML(http.StatusOK, "ch02/arr_struct.html", arr_struct)
}

// Map data render
func MapHandler(ctx *gin.Context) {
	map_data := map[string]interface{}{
		"name": "tom",
		"age":  18,
	}

	ctx.HTML(http.StatusOK, "ch02/map.html", map_data)

}

// Map and struct render
func MapAndStruct(ctx *gin.Context) {
	map_struct_data := map[string]UserInfo{
		"user": {ID: 3, Name: "tom", Age: 19, Addr: "xxxx"},
	}

	ctx.HTML(http.StatusOK, "ch02/map_struct.html", map_struct_data)
}

// Slice data render

// 1. 带参数的路由：路径中直接加上参数值
// 使用占位符: ，必须得指定这个路径
func Param1(ctx *gin.Context) {

	name := ctx.Param("name")

	ctx.String(http.StatusOK, "hello, %s", name)
}

// 使用占位符*，可以不用匹配这个路径
func Param2(ctx *gin.Context) {

	name := ctx.Param("name")

	ctx.String(http.StatusOK, "hello, %s", name)
}

// 二、带参数的路由：路径中使用 param name
func GetQueryData(ctx *gin.Context) {
	id := ctx.Query("id")

	name := ctx.DefaultQuery("name", "mark")

	ctx.String(http.StatusOK, "hello, %s, %s", id, name)
}

func GetQueryArrData(ctx *gin.Context) {
	ids := ctx.QueryArray("ids")

	ctx.String(http.StatusOK, "hello, %v", ids)
}

func GetQueryMapData(ctx *gin.Context) {
	user := ctx.QueryMap("user")

	ctx.String(http.StatusOK, "hello, %v", user)
}

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

// Parameter Bind
// go to user add page
func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/user_add4.html", nil)
}

// user add
func DoUserAdd4(ctx *gin.Context) {
	var user_info UserInfo

	ctx.ShouldBind(&user_info)
	// ctx.ShouldBindWith()

	fmt.Println(user_info)

	ctx.String(http.StatusOK, "user add ok!")
}
