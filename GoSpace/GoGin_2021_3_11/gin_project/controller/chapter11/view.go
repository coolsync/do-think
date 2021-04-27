package chapter11

import (
	"ginproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Books struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Url string `json:"url"`
}

func ApiAxios(ctx *gin.Context) {

	user := models.User{
		ID:   1,
		Name: "paul",
		Age:  18,
	}

	// slice
	arrs := []string{"a", "b", "c", "d", "e", "f"}

	// struct slice
	arrs_struct := []models.User{
		{
			ID: 2, Name: "paul2", Age: 19,
		},
		{
			ID: 3, Name: "paul3", Age: 20,
		},
	}

	// map slice
	map_struct := map[string]models.User{
		"user": {
			ID: 4, Name: "paul4", Age: 19,
		},
	}
	ctx.JSON(
		http.StatusOK,
		gin.H{
			"code":   200,
			"msg":    "提交成功",
			"user":   user,
			"arrs":   arrs,
			"arrs_s": arrs_struct,
			"map_s":  map_struct,
		})
}

func GetBooks(ctx *gin.Context) {
	books := []Books {
		{
			ID: 1, Name: "Go by Example 中文", Url: "https://books.studygolang.com/gobyexample/",
		},
		{
			ID: 2, Name: "Go RPC 开发指南", Url: "https://books.studygolang.com/go-rpc-programming-guide/",
		},
		{
			ID: 3, Name: "Go语言高级编程", Url: "https://books.studygolang.com/advanced-go-programming-book/",
		},
		{
			ID: 4, Name: "Mastering_Go_ZH_CN", Url: "https://books.studygolang.com/Mastering_Go_ZH_CN/",
		},
	}

	ctx.JSON(http.StatusOK, gin.H{
		"books": books,
	})
}