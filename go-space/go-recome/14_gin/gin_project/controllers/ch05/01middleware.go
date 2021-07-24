package ch05

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare1(ctx *gin.Context) {
	time_now := time.Now()

	fmt.Println("This is custom middleware 1, Start")
	ctx.Next()
	fmt.Println("This is custom middleware 1, End")

	time_count := time.Since(time_now)
	fmt.Printf("total cost time: %v\n", time_count.Seconds())
}

func MiddleWare2() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("---- This is custom middleware 2, Start")
		// b := true
		// if b {
		// 	ctx.Abort() // ctx.Abort()方法的作用 终止调用整个 middle chain, gin back all content is not visit
		// }

		fmt.Println("---- MiddleWare2 sleepping ...")
		time.Sleep(time.Second * 2)

		ctx.Next()
		fmt.Println("---- This is custom middleware 2, End")
	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("---- ---- This is custom middleware 3, Start")
	ctx.Next()
	fmt.Println("---- ---- This is custom middleware 3, End")
}
