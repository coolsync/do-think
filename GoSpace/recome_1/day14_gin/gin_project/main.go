package main

import (
	"gin2/ch01"
	"gin2/ch02"
	"gin2/ch03"
	"gin2/ch04"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	// Create router
	router := gin.Default()

	// Set consume tpl func
	router.SetFuncMap(template.FuncMap{
		"add_num":   ch03.AddNum,
		"str_len":   ch03.SubStr,
		"safe_html": ch03.SafeHTML,
	})

	// Custom validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("len6_valid", ch04.Len6Validator) //
	}

	// Use tmpl Regular, specify 多级目录
	router.LoadHTMLGlob("templates/**/*")

	// Specify static file dir
	router.Static("/static", "static")
	// router.StaticFS("/static", http.Dir("static"))

	// Ch01
	router.GET("/", ch01.Hello)

	// Ch02 Tmpls Render
	router.GET("/user", ch02.User)
	router.GET("/user_info", ch02.UserInfoStruct)
	router.GET("/arr", ch02.ArrayHandler)
	router.GET("/arr_struct", ch02.ArrayAndStruct)
	router.GET("/map", ch02.MapHandler)
	router.GET("/map_struct", ch02.MapAndStruct)

	router.GET("/param1/:name", ch02.Param1)
	router.GET("/param2/*name", ch02.Param2)

	router.GET("/query", ch02.GetQueryData)
	router.GET("/query_arr", ch02.GetQueryArrData)
	router.GET("/query_map", ch02.GetQueryMapData)

	// PostForm
	router.GET("/to_user_add", ch02.ToUserAdd)
	router.POST("/do_user_add", ch02.DoUserAdd)

	// DefaultPostForm, DefaultQuery, ctx.QueryArray, ctx.QueryMap
	router.GET("/to_user_add2", ch02.ToUserAdd2)
	router.POST("/do_user_add2", ch02.DoUserAdd2)

	// Ajax Req
	router.GET("/to_user_add3", ch02.ToUserAdd3)
	router.POST("/do_user_add3", ch02.DoUserAdd3)

	// Parameter Bind
	router.GET("/to_user_add4", ch02.ToUserAdd4)
	router.POST("/do_user_add4", ch02.DoUserAdd4)

	// Upload File
	router.GET("/to_upload1", ch02.ToUpload1) // single file upload
	router.POST("/do_upload1", ch02.DoUpload1)

	router.GET("/to_upload2", ch02.ToUpload2) // multiple file upload
	router.POST("/do_upload2", ch02.DoUpload2)

	// Ajax Upload File
	router.GET("/to_upload3", ch02.ToUpload3) // ajax single file upload
	router.POST("/do_upload3", ch02.DoUpload3)

	router.GET("/to_upload4", ch02.ToUpload4) // ajax multiple file upload
	router.POST("/do_upload4", ch02.DoUpload4)

	// Output Other data format type
	router.GET("/out_json", ch02.OutJson)
	router.GET("/out_ascii_json", ch02.OutAsciiJSON)
	router.GET("/out_jsonp", ch02.OutJSONP)
	router.GET("/out_pure_json", ch02.OutPureJSON)
	router.GET("/out_secure_json", ch02.OutSecureJSON)

	router.GET("/out_xml", ch02.OutXML)
	router.GET("/out_yaml", ch02.OutYaml)
	router.GET("/out_protobuf", ch02.OutProto)

	// Redirect
	router.GET("/redirect_a", ch02.RedirectA)
	router.GET("/redirect_b", ch02.RedirectB)

	// Ch03
	// tpl syntax
	router.GET("/tpl_syntax1", ch03.TplSyntax1)
	router.GET("/tpl_syntax2", ch03.TplSyntax2)
	// tpl func
	router.GET("/tpl_func1", ch03.TplFunc1)
	router.GET("/tpl_func2", ch03.TplFunc2)
	router.GET("/consume_tpl_func", ch03.ConsumeTplFunc)

	// Ch04
	// Data Binding
	router.GET("/to_bind_form", ch04.ToBindForm)
	router.POST("/do_bind_form", ch04.DoBindForm)

	router.GET("/to_bind_json", ch04.ToBindJson)
	router.POST("/do_bind_json", ch04.DoBindJson)

	router.GET("/bind_query", ch04.GetQueryData)
	router.GET("/bind_uri/:name/:age/:addr", ch04.BindUri)

	// Validator
	router.GET("/to_valid", ch04.ToValidData)
	router.POST("/do_valid", ch04.DoValidData)

	// Beego Validator
	router.GET("/to_beego_validator", ch04.ToBeegoValidator)
	router.POST("/do_beego_validator", ch04.DoBeegoValidator)

	// Listen port
	// router.Run(":8090")

	// Custom HTTP Configuration
	// http.ListenAndServe(":8090", router)

	s := &http.Server{
		Addr:           ":8090",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1 * 2^20
	}
	s.ListenAndServe()
}