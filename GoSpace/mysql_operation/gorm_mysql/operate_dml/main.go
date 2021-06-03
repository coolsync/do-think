package main

import (
	"comegorm/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func DMLHandler1() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Create
	// db.Create(&models.User{Name: "bob", Age: 30, Addr: "xxx", Pic: "/static/upload/pic.jpg"})

	// Query
	var user1 models.User
	db.First(&user1, 2) // id = 2
	// db.First(&user1, "name=?", "bob")
	fmt.Println("user1: ", user1)

	// Update
	var user2 models.User
	db.First(&user2, 2)
	user2.Name = "paul"
	user2.Age = 20
	db.Save(&user2)

	db.Model(&user2).Update("addr", "paul-yyyy")
	db.Model(&user2).Update("phone", "12345678")

	// Del
	var user3 models.User
	db.First(&user3, 4)
	db.Delete(&user3)

	db.Where("name", "mark").Delete(&models.User{})
}
