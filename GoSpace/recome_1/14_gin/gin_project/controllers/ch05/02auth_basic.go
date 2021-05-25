package ch05

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// simulate some private data
// var map_data map[string]interface{} = map[string]interface{}{
// 	"bob":   gin.H{"age": 18, "addr": "bob--xxx"},
// 	"paul":  gin.H{"age": 19, "addr": "paul--xxx"},
// 	"jerry": gin.H{"age": 20, "addr": "jerry--xxx"},
// }

var secrets_data = gin.H{
	"bob":   gin.H{"age": 18, "addr": "bob--xxx"},
	"paul":  gin.H{"age": 19, "addr": "paul--xxx"},
	"jerry": gin.H{"age": 20, "addr": "jerry--xxx"},
}

func AuthBasicHandler(ctx *gin.Context) {
	// get query params val
	user_name := ctx.Query("user_name")

	// query user is not exists
	user_info, ok := secrets_data[user_name]

	if ok {
		ctx.JSON(http.StatusOK, gin.H{
			"user": user_name,
			"data": user_info,
		})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{
			"user": user_name,
			"data": "user is not admin",
		})
	}
}

func WrapFHandler(w http.ResponseWriter, r *http.Request) {

}
