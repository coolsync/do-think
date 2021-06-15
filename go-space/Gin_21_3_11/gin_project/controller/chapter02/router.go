package chapter02

import (
	"ginproject/controller/chapter05"

	"github.com/gin-gonic/gin"
)

func Router(ch02 *gin.RouterGroup) {
	ch02.GET("/user", User)
	ch02.GET("/user_info", chapter05.MiddleWare3, UserInfoStruct)	// 局部使用中间件
	ch02.GET("/arr", ArrayHandler)
	ch02.GET("/arrstruct", ArrayAndStructHandler)
	ch02.GET("/map", MapHandler)
	ch02.GET("/map_struct", MapAndStructHandler)
	ch02.GET("/param1/:id", Param1Handler)
	ch02.GET("/param2/*id", Param2Handler) // 会把 前面的  / plus

	// Get the data submitted by the form， Ajax 交互
	ch02.GET("/to_user_add3", ToUserAdd3)
	ch02.POST("/do_user_add3", DoUserAdd3)

	// ShouldBind
	ch02.GET("/to_user_add4", ToUserAdd4)
	ch02.POST("/do_user_add4", DoUserAdd4)

	// Form upload single file
	ch02.GET("/to_upload1", ToUpload1)
	ch02.POST("/do_upload1", DoUpload1)

	// Form upload multiple file
	ch02.GET("/to_upload2", ToUpload2)
	ch02.POST("/do_upload2", DoUpload2)

	// Ajax form upload single file
	ch02.GET("/to_ajax_upload3", ToUploadFile3)
	ch02.POST("/do_ajax_upload3", DoUploadFile3)

	// Ajax form upload multiple file
	ch02.GET("/to_ajax_upload4", ToUploadFile4)
	ch02.POST("/do_ajax_upload4", DoUploadFile4)

	// other output, json ascii-json
	ch02.GET("/output_json", OutJson)
	ch02.GET("/output_ascii_json", OutAsciiJson)
	ch02.GET("/output_jsonp", OutJsonP)
	ch02.GET("/output_purejson", OutPureJson)
	ch02.GET("/output_securejson", OutSecureJson)
	ch02.GET("/output_xml", OutXml)
	ch02.GET("/output_yaml", OutYaml)
	ch02.GET("/output_protobuf", OutProtobuf)

	// Router Redirect
	ch02.GET("/redirect_a", RedirectA)
	ch02.GET("/redirect_b", RedirectB)
}
