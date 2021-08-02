package chapter04

import (
	"fmt"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// Binding
// type Article struct {
// 	Id      int    `form:"-"`
// 	Title   string `form:"title" binding:"len_valid"`	// custom validator
// 	Content string `form:"content" binding:"min=4"`
// 	Desc    string `form:"desc"`
// }

// func ToBindValid(ctx *gin.Context) {
// 	ctx.HTML(http.StatusOK, "chapter04/bind_valid.html", nil)
// }

// func DoBindValid(ctx *gin.Context) {
// 	var article Article

// 	err := ctx.ShouldBind(&article)
// 	fmt.Println(err)
// 	fmt.Println(article)

// 	ctx.String(http.StatusOK, "OK")
// }

// 1 Define validator
// 2 register validator
// 3 in struct use
var Len6Validator validator.Func = func(fl validator.FieldLevel) bool {

	data := fl.Field().Interface().(string)

	fmt.Println(data)

	if len(data) > 6 {
		fmt.Println("false no pass")
		return false
	}

	fmt.Println("true ok pass")
	return true
}

// Valid for Beego
type Article struct {
	Id      int    `form:"-"`
	Title   string `form:"title" valid:"Required"`
	Content string `form:"content" valid:"Length(6)"`
	Desc    string `form:"desc" valid:"MinSize(3)"`
	Email   string `form:"email" valid:"Email"`
}

func ToBindValid(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "chapter04/bind_valid.html", nil)
}

func DoBindValid(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	fmt.Println(err)
	fmt.Println(article)

	// init beego valid
	valid := validation.Validation{}
	fmt.Println("valid: ", valid)

	//重写错误信息：validation.SetDefaultMessage(map)
	// MessageTmpls store commond validate template
	var MessageTmpls = map[string]string{
		// "Required":     "Can not be empty",
		"Required": "不能为空", // modify "Required"
		"Min":      "Minimum is %d",
		"Max":      "Maximum is %d",
		"Range":    "Range is %d to %d",
		// "MinSize":      "Minimum size is %d",
		"MinSize": "字符串最小长度为 %d",
		"MaxSize": "Maximum size is %d",
		// "Length":       "Required length is %d",
		"Length":       "字符串固定长度为 %d",
		"Alpha":        "Must be valid alpha characters",
		"Numeric":      "Must be valid numeric characters",
		"AlphaNumeric": "Must be valid alpha or numeric characters",
		"Match":        "Must match %s",
		"NoMatch":      "Must not match %s",
		"AlphaDash":    "Must be valid alpha or numeric or dash(-_) characters",
		// "Email":        "Must be a valid email address",
		"Email":   "必须是有效邮箱的",
		"IP":      "Must be a valid ip address",
		"Base64":  "Must be valid base64 characters",
		"Mobile":  "Must be valid mobile number",
		"Tel":     "Must be valid telephone number",
		"Phone":   "Must be valid telephone or mobile phone number",
		"ZipCode": "Must be valid zipcode",
	}

	validation.SetDefaultMessage(MessageTmpls)

	key_mapping := map[string]string{
		"Title.Required.": "标题",
		"Content.Length.": "内容",
		"Desc.MinSize.":   "描述",
		"Email.Email.":    "邮箱",
	}

	bl, err1 := valid.Valid(&article)
	fmt.Println("err1: ", err1)

	if !bl { // valid err
		for _, err := range valid.Errors {
			fmt.Println(err.Key)
			fmt.Println(err.Message)
			ctx.String(http.StatusOK, key_mapping[err.Key]+" "+err.Message+"\n")
		}
	} else {
		ctx.String(http.StatusOK, "OK!")
	}
}
