package chapter11

import (
	"ginproject/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

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
		"user": models.User{
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
