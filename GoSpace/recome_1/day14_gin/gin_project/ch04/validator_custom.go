package ch04

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// data verification
// Binding
type Article struct {
	ID          int        `form:"id"`
	Title       string     `form:"title" binding:"omitempty,len6_valid"`
	Content     string     `form:"content" binding:"required,min=5"`
	Description string     `form:"desc" binding:"max=10"`
	Uid         []UserInfo `binding:"omitempty,dive,min=10"` // dive nested verify，split nested
	//Name [][]int `binding:"min=10,dive,max=20,dive,required"`
	Name [][]int `binding:"dive,max=20,dive,required"`
}

func ToValidData(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "ch04/validator.html", nil)
}

func DoValidData(ctx *gin.Context) {
	var article Article

	err := ctx.ShouldBind(&article)
	if err != nil {
		fmt.Println(err)
		ctx.String(http.StatusNotFound, "get article failed: %v\n", article)
		return
	}
	ctx.String(http.StatusOK, "Ok: %v\n", article)
}

// Custom validator
// 1 define validator
var Len6Validator validator.Func = func(fl validator.FieldLevel) bool {
	data := fl.Field().Interface().(string)

	fmt.Println(data, len(data))

	// if len(data) > 6 { // validator > 6, 则不通过
	// 	return true
	// }
	// return false

	return len(data) > 6
}

// 2 in main, register validator
// if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
// 	v.RegisterValidation("len6_valid", ch04.Len6Validator)	//
// }

// 3 use in struct
// Title       string     `form:"title" binding:"omitempty,len6_valid"`
