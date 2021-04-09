package chapter06

import (
	"fmt"
	"ginproject/models"

	"github.com/gin-gonic/gin"
	"ginproject/data_source"
)
// 

func GormTest(ctx *gin.Context) {
	// var user models.User
	user := models.User{Name: "bob", Age: 30}
	data_source.Db.Create(&user)
	fmt.Println(user)
	
}