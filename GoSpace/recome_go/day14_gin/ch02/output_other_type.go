package ch02

import (
	"net/http"

	user "gin2/proto"

	"github.com/gin-gonic/gin"
)

func OutJson(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})

	// will output:
	// {"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"提交成功","tag":"\u003cbr\u003e"}
}

// Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII characters.
func OutAsciiJSON(ctx *gin.Context) {
	map_data := map[string]interface{}{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	}

	ctx.AsciiJSON(http.StatusOK, map_data)

	// output:
	// {"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"\u63d0\u4ea4\u6210\u529f","tag":"\u003cbr\u003e"}
}

// Using JSONP to request data from a server in a different domain.
// Add callback to response body if the query parameter callback exists.
func OutJSONP(ctx *gin.Context) {
	ctx.JSONP(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})

	// output:
	// {"code":200,"html":"\u003cb\u003eHello, world!\u003c/b\u003e","msg":"提交成功","tag":"\u003cbr\u003e"}
}

// Normally, JSON replaces special HTML characters with their unicode entities, e.g. < becomes \u003c. If you want to encode such characters literally, you can use PureJSON instead. This feature is unavailable in Go 1.6 and lower.
func OutPureJSON(ctx *gin.Context) {
	ctx.PureJSON(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"msg":  "提交成功",
		"html": "<b>Hello, world!</b>",
	})
	// output:
	// {"code":200,"html":"<b>Hello, world!</b>","msg":"提交成功","tag":"<br>"}
}

// Using SecureJSON to prevent json hijacking. Default prepends "while(1)," to response body if the given struct is array values.
func OutSecureJSON(ctx *gin.Context) {
	names := []string{"bob", "mark", "paul"}
	ctx.SecureJSON(http.StatusOK, names)

	// while(1);["bob","mark","paul"], not parse json data
}

func OutYaml(ctx *gin.Context) {
	ctx.YAML(http.StatusOK, gin.H{
		"code": 200,
		"tag":  "<br>",
		"user": gin.H{"name": "pual", "age": 18},
		"html": "<b>Hello, world!</b>",
	})
}

func OutProtoBuf(ctx *gin.Context) {
	// The specific definition of protobuf is written in the testdata/protoexample file.
	data := &user.UserResponse{
		Id:   1,
		Name: "paul",
		Age:  18,
	}
	// Note that data becomes binary data in the response
	// Will output protoexample.Test protobuf serialized data
	ctx.ProtoBuf(http.StatusOK, data)
}
