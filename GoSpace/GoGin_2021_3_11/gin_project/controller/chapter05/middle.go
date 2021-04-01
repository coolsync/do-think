package chapter05

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func MiddleWare1(ctx *gin.Context) {
	time_start := time.Now()
	fmt.Println("custom middleware 1 -- start")
	ctx.Next()

	time_count := time.Since(time_start)
	fmt.Println(time_count)
	fmt.Println("custom middleware 1 -- end")
}

func MiddleWare2() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("custom middleware 2 -- start")
		// b := true
		// if b {
		// 	c.Abort() // abort middleware
		// }
		time.Sleep(time.Second * 2)
		c.Next()
		fmt.Println("custom middleware 2 -- end")

	}
}

func MiddleWare3(ctx *gin.Context) {
	fmt.Println("custom middleware 3 -- start")
	time.Sleep(time.Second * 2)
	ctx.Next()
	fmt.Println("custom middleware 3 -- end")
}
