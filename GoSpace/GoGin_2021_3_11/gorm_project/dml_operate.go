package main

import (
	"fmt"
	"gorm_project/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	/*
			 // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
		  	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		  	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	*/
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{})

	// Create
	// db.Create(&models.User{Name:"bob", Age: 30, Addr: "xxx", Pic: "/static/upload/pic.jpg"})

	var user models.User

	// Query
	// db.First(&user, 1)	// 1 is id
	// db.First(&user, "name=?", "bob")
	// fmt.Println(user)

	// // Update
	// db.First(&user, 2)
	// // 1
	// user.Name = "paul"
	// user.Age = 20
	// db.Save(&user)

	// // 2
	// db.Model(&user).Update("addr", "pual-xxxx")
	// db.Model(&user).Update("phone", "12345678")

	// // 3
	// db.Model(&user).Updates(models.User{Name: "jerry", Addr: "jerry-xxxx"})

	// Delete
	// db.Delete(&user, 2)
	db.Where("name", "bob").Delete(&user)

	fmt.Println("DML OK")
}
