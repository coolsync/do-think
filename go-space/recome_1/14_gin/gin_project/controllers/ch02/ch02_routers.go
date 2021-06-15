package ch02

import (
	"gin2/controllers/ch05"

	"github.com/gin-gonic/gin"
)

func Routers(ch02_routers *gin.RouterGroup) {
	// Ch02 Tmpls Render
	ch02_routers.GET("/user", User)

	ch02_routers.GET("/user_info", ch05.MiddleWare3, UserInfoStruct) // Partial use of middleware
	
	ch02_routers.GET("/arr", ArrayHandler)
	ch02_routers.GET("/arr_struct", ArrayAndStruct)
	ch02_routers.GET("/map", MapHandler)
	ch02_routers.GET("/map_struct", MapAndStruct)

	ch02_routers.GET("/param1/:name", Param1)
	ch02_routers.GET("/param2/*name", Param2)

	ch02_routers.GET("/query", GetQueryData)
	ch02_routers.GET("/query_arr", GetQueryArrData)
	ch02_routers.GET("/query_map", GetQueryMapData)

	// PostForm
	ch02_routers.GET("/to_user_add", ToUserAdd)
	ch02_routers.POST("/do_user_add", DoUserAdd)

	// DefaultPostForm, DefaultQuery, ctx.QueryArray, ctx.QueryMap
	ch02_routers.GET("/to_user_add2", ToUserAdd2)
	ch02_routers.POST("/do_user_add2", DoUserAdd2)

	// Ajax Req
	ch02_routers.GET("/to_user_add3", ToUserAdd3)
	ch02_routers.POST("/do_user_add3", DoUserAdd3)

	// Parameter Bind
	ch02_routers.GET("/to_user_add4", ToUserAdd4)
	ch02_routers.POST("/do_user_add4", DoUserAdd4)

	// Upload File
	ch02_routers.GET("/to_upload1", ToUpload1) // single file upload
	ch02_routers.POST("/do_upload1", DoUpload1)

	ch02_routers.GET("/to_upload2", ToUpload2) // multiple file upload
	ch02_routers.POST("/do_upload2", DoUpload2)

	// Ajax Upload File
	ch02_routers.GET("/to_upload3", ToUpload3) // ajax single file upload
	ch02_routers.POST("/do_upload3", DoUpload3)

	ch02_routers.GET("/to_upload4", ToUpload4) // ajax multiple file upload
	ch02_routers.POST("/do_upload4", DoUpload4)

	// Output Other data format type
	ch02_routers.GET("/out_json", OutJson)
	ch02_routers.GET("/out_ascii_json", OutAsciiJSON)
	ch02_routers.GET("/out_jsonp", OutJSONP)
	ch02_routers.GET("/out_pure_json", OutPureJSON)
	ch02_routers.GET("/out_secure_json", OutSecureJSON)

	ch02_routers.GET("/out_xml", OutXML)
	ch02_routers.GET("/out_yaml", OutYaml)
	ch02_routers.GET("/out_protobuf", OutProto)

	// Redirect
	ch02_routers.GET("/redirect_a", RedirectA)
	ch02_routers.GET("/redirect_b", RedirectB)

}
