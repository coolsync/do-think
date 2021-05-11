package ch02

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	ID   int
	Name string
	Age  int
	Addr string
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
