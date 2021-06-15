package chapter02

import (
	"fmt"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
)

func OutJson(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

// 使用 AsciiJSON 方法可以生成只包含 ASCII 字符的 JSON 格式数据，对于非 ASCII 字符会进行转义
func OutAsciiJson(ctx *gin.Context) {
	ctx.AsciiJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

// Get data across domains, 跨域获取数据
func OutJsonP(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

//
func OutPureJson(ctx *gin.Context) {
	ctx.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
	})
}

func OutSecureJson(ctx *gin.Context) {
	// names := []string{"lena", "austin", "foo"}
	// 将输出：while(1);["lena","austin","foo"]
	// ctx.SecureJSON(http.StatusOK, names)

	ctx.SecureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

func OutXml(ctx *gin.Context) {
	ctx.XML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
}

func OutYaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"user": gin.H{"name": "zhiliao", "age": 18},
		"html": "<b>Hello, world!</b>",
	})
}

func OutProtobuf(ctx *gin.Context) {
	user_data := user.User{
		Name: "bob",
		// Age:  int32(18),
	}

	ctx.ProtoBuf(http.StatusOK, &user_data)
}

// Redirect
func RedirectA(ctx *gin.Context) {
	fmt.Println("这是 路由 A！")

	ctx.Redirect(http.StatusFound, "/redirect_b") // 302
	// ctx.Redirect(http.StatusFound, "https://cn.bing.com/search?form=MOZLBR&pc=MOZI&q=gnome+extension")
}

func RedirectB(ctx *gin.Context) {
	fmt.Println("这是 路由 B！")

	ctx.String(http.StatusOK, "这是 路由 B！")
}


