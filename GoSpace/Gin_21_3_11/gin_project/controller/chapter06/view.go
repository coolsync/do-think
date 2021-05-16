package chapter06

import (
	"fmt"
	"ginproject/models"

	"ginproject/data_source"

	"github.com/gin-gonic/gin"
)

func GormTest(ctx *gin.Context) {
	// var user models.User
	user := models.User{Name: "bob", Age: 30}
	data_source.Db.Create(&user)
	fmt.Println(user)
}
