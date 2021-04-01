package chapter02

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
	Id   int    `form:"id"`
	Name string `form:"name"`
	Age  int    `form:"age"`
	Addr string `form:"addr"`
}

// ajax add user
func ToUserAdd3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_add3.html", nil)
}

func DoUserAdd3(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	age := ctx.PostForm("age")

	fmt.Println(username)
	fmt.Println(password)
	fmt.Println(age)

	map_data := map[string]interface{}{
		"code": 200,
		"msg":  "OK",
	}

	ctx.JSON(http.StatusOK, map_data)
}

// add user page handler
func ToUserAdd4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user_add4.html", nil)
}

func DoUserAdd4(ctx *gin.Context) {

	var user_info UserInfo

	err := ctx.ShouldBind(&user_info) // Corresponding struct
	fmt.Println(err)
	fmt.Println(user_info)

	ctx.String(http.StatusOK, "完成绑定")
}

// Form upload single file
func ToUpload1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/test_upload1.html", nil)
}

func DoUpload1(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	fmt.Println(file.Filename)

	time_unix_int := time.Now().UnixNano()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)
	dst := "upload/" + time_unix_str + file.Filename
	ctx.SaveUploadedFile(file, dst)

	ctx.String(http.StatusOK, "file upload ok!")
}

// Form upload multiple file
func ToUpload2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/test_upload2.html", nil)
}

func DoUpload2(ctx *gin.Context) {
	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, file := range files {
		time_unix_int := time.Now().UnixNano()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		dst := "upload/" + time_unix_str + file.Filename
		ctx.SaveUploadedFile(file, dst)
	}

	ctx.String(http.StatusOK, "file upload ok!")
}

// Ajax form upload single file
func ToUploadFile3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/ajax_upload3.html", nil)
}
func DoUploadFile3(ctx *gin.Context) {

	name := ctx.PostForm("name")
	fmt.Println(name)

	file, _ := ctx.FormFile("file")

	time_unix_int := time.Now().UnixNano()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)
	dst := "upload/" + time_unix_str + file.Filename
	ctx.SaveUploadedFile(file, dst)

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// Ajax form upload multiple file
func ToUploadFile4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter02/ajax_upload4.html", nil)
}

func DoUploadFile4(ctx *gin.Context) {
	name := ctx.PostForm("name")
	fmt.Println(name)

	form, _ := ctx.MultipartForm()
	files := form.File["file"]

	for _, f := range files {
		fmt.Println(f.Filename)
		time_unix_int := time.Now().UnixNano()
		time_unix_str := strconv.FormatInt(time_unix_int, 10)
		dst := "upload/" + time_unix_str + f.Filename
		ctx.SaveUploadedFile(f, dst)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}
