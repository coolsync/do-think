package main

import (
	"ginproject/controller/chapter03"
	"ginproject/controller/chapter04"
	"ginproject/controller/chapter05"
	"ginproject/routers"
	"html/template"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	_ "ginproject/data_source"
	_ "ginproject/logs_source"
)

func main() {
	router := gin.Default()

	// router := gin.New()
	// router.Use(gin.Logger(), gin.Recovery())

	// // Gin 自带 logger
	// f, _ := os.Create("./gin_log")
	// gin.DefaultWriter = io.MultiWriter(f)

	// Global Middleware
	// router.Use(chapter05.MiddleWare1)
	router.Use(chapter05.MiddleWare2())
	// router.Use(chapter05.MiddleWare3)

	// Costom tmpl func
	router.SetFuncMap(template.FuncMap{
		"add":    chapter03.AddNum,
		"substr": chapter03.SubStr,
		"safe":   chapter03.Safe,
	})

	// Bind validator
	// register validator
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("len_valid", chapter04.Len6Validator)
	}

	// Router group
	routers.Routers(router)

	router.LoadHTMLGlob("template/**/*")
	// router.LoadHTMLFiles("index.html, news.html")

	// Load static source
	router.Static("/static", "static")

	// router.Run(":9000")

	// Custom HTTP configuration， 自定义 HTTP 配置
	// http.ListenAndServe(":9000", router)

	server := &http.Server{
		Addr:           ":9000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()
}
