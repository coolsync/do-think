package ch04

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	// "github.com/astaxie/beego/validation"
	"github.com/beego/beego/v2/core/validation"
	"github.com/gin-gonic/gin"
)

type user struct {
	Id   int    `form:"id"`
	Name string `form:"name" valid:"Required;Length(5)"`
	// Age   int    `form:"age" valid:"Required;Min(18);Max(49)"`
	Age   int    `form:"age" valid:"Required;Range(18, 49)"`
	Email string `form:"email" valid:"Email;MaxSize(100)"`
	Phone string `form:"phone" valid:"Phone"`
	IP    string `form:"ip" valid:"IP"`
}

// If your struct implemented interface `validation.ValidFormer`
// When all tests in StructTag succeed, it will execute Valid function for custom validation
func (u *user) Valid(v *validation.Validation) {
	if strings.Index(u.Name, "admin") != -1 {
		// Set error messages of Name by SetError and HasErrors will return true
		v.SetError("Name", "Can't contain 'admin' in Name")
	}
	fmt.Println(1234)
}

// show page
func ToBeegoValidator(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/beego_validator.html", nil)
}

// handler page
func DoBeegoValidator(ctx *gin.Context) {
	var u *user
	// bind data to sturct
	err := ctx.ShouldBind(&u)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusNotFound, "server handler failed")
		return
	}

	// init beego validator
	valid := validation.Validation{}

	key_mapping := map[string]interface{}{
		"Name.Required.": "姓名",
		"Name.Length.":   "姓名Length",
		"Age.Min.":       "年龄",
		"Age.Max.":       "年龄",
		"Age.Range.":     "年龄",
		"Email.Email.":   "邮箱",
		"Phone.Phone.":   "手机",
		"IP.IP.":         "ip addr",
	}

	// modify default error messsages
	messages_mapping := map[string]string{
		// "Required":     "Can not be empty",
		// "Min":          "Minimum is %d",
		// "Max":          "Maximum is %d",
		// "Range":        "Range is %d to %d",
		// "Length":       "Required length is %d",
		"Required": "不能为空",
		"Min":      "最小值 %d",
		"Max":      "最大值 is %d",
		"Range":    "范围 从 %d 到 %d",
		"Length":   "所需长度为  %d",
		// "Email":        "Must be a valid email address",
		// "IP":           "Must be a valid ip address",
		// "Phone":        "Must be valid telephone or mobile phone number",
		"Email": "必须是一个有效的E-mail地址",
		"IP":    "必须是有效的IP地址 ",
		"Phone": "必须是有效的电话或手机号码",
	}
	validation.SetDefaultMessage(messages_mapping)

	// bind sturct
	b, err := valid.Valid(&u)
	// u.Valid(&valid)
	if err != nil {
		log.Println(err)
		ctx.String(http.StatusNotFound, "server handler failed")
		return
	}

	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key)
			log.Println(err.Message)
			ctx.String(http.StatusBadRequest, "%s: %s\n", err.Key, err.Message)
			ctx.String(http.StatusBadRequest, "%s: %s\n", key_mapping[err.Key].(string), err.Message)
			return
		}
	}

	// if valid.HasErrors() {
	// 	// extract error
	// 	for _, err := range valid.Errors {
	// 		log.Println(err.Key, err.Message)
	// 		return
	// 	}
	// }

	ctx.String(http.StatusOK, "OK: %v\n", u)
}
