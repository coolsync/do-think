package main

import (
	"gin2/controllers/ch03"
	"gin2/controllers/ch04"
	"gin2/routers"
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
	// router := gin.New()
	// router.Use(gin.Logger(), gin.Recovery()) // built middleware

	// router.Use(ch05.MiddleWare1) // global middle ware
	// router.Use(ch05.MiddleWare2())
	// router.Use(ch05.MiddleWare3)
	// router.Use(ch05.MiddleWare1, ch05.MiddleWare2(), ch05.MiddleWare3)

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

	routers.AllRouters(router) // Router Group

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
