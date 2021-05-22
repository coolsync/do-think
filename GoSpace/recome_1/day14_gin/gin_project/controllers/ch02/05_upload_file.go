package ch02

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// Form Upload Single File
func ToUpload1(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/upload_file1.html", nil)
}

func DoUpload1(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")

	fmt.Println(file.Filename)

	time_unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(time_unix_int, 10)

	dst_path := "upload/" + time_unix_str + "_" + file.Filename
	ctx.SaveUploadedFile(file, dst_path)

	ctx.String(http.StatusOK, "Upload File Ok!")
}

// Form Upload Multiple File
func ToUpload2(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/upload_file2.html", nil)
}

func DoUpload2(ctx *gin.Context) {
	// get form obj, operate mul file
	form, _ := ctx.MultipartForm()
	// get files
	files := form.File["file"]
	// for files, get file obj
	// save file

	for _, file := range files {
		fmt.Println(file.Filename)
		unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(unix_int, 10)
		dst_path := "upload/" + time_unix_str + "_" + file.Filename
		ctx.SaveUploadedFile(file, dst_path)
	}

	ctx.String(http.StatusOK, "Upload File Ok!")
}

// Ajax Upload Single File
func ToUpload3(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/upload_ajax1.html", nil)
}

func DoUpload3(ctx *gin.Context) {
	name := ctx.PostForm("name")

	fmt.Println(name)

	file, _ := ctx.FormFile("file")

	fmt.Println(file.Filename)

	unix_int := time.Now().Unix()
	time_unix_str := strconv.FormatInt(unix_int, 10)

	dst_path := "upload/" + time_unix_str + "_" + file.Filename
	ctx.SaveUploadedFile(file, dst_path)

	// ctx.String(http.StatusOK, "Upload )
	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}

// Ajax Upload Multiple File
func ToUpload4(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch02/upload_ajax2.html", nil)
}

func DoUpload4(ctx *gin.Context) {
	name := ctx.PostForm("name")

	fmt.Println(name)

	form, err := ctx.MultipartForm()
	if err != nil {
		fmt.Println(err)
	}

	files := form.File["file"]

	// for files, handle multiple file
	for _, file := range files {
		fmt.Println(file.Filename)

		unix_int := time.Now().Unix()
		time_unix_str := strconv.FormatInt(unix_int, 10)

		dst_path := "upload/" + time_unix_str + "_" + file.Filename // Splicing path
		_ = ctx.SaveUploadedFile(file, dst_path)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "OK",
	})
}
