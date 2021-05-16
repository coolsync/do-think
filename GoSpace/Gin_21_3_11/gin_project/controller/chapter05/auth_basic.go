package chapter05

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// set private data
var map_data map[string]interface{} = map[string]interface{}{
	"bob":   gin.H{"age": "18", "addr": "bob--xxx"},
	"paul":  gin.H{"age": "19", "addr": "paul--xxx"},
	"jerry": gin.H{"age": "20", "addr": "jerry--xxx"},
}

// get private data
func AuthBasicTest(ctx *gin.Context) {

	user_name := ctx.Query("user_name")

	data, ok := map_data[user_name]

	if ok {
		ctx.JSON(http.StatusOK, gin.H{"user_name": user_name, "data": data})
	} else {
		ctx.JSON(http.StatusNotFound, gin.H{"user_name": user_name, "data": "no 权限"})
	}
}

func WrapFDisc(w http.ResponseWriter, r *http.Request) {
	
}
