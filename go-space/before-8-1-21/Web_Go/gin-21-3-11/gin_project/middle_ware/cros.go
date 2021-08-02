package middle_ware

import "github.com/gin-gonic/gin"

func CrosMiddleWare(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
}
