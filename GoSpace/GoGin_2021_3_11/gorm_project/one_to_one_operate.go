package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var P = fmt.Println

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

	// // one to one
	db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})

	// Insert a recode

	// user_profile := relate_tables.UserProfile {
	// 	Pic:   "1.jpg",
	// 	CPic:  "2.jpg",
	// 	Phone: "12345678909",
	// 	User: relate_tables.User{
	// 		Name: "bob",
	// 		Age:  30,
	// 		Addr: "guangdong shengzheng",
	// 	},
	// }
	// db.Create(&user_profile)

	// 1 Association Query
	var user_profile relate_tables.UserProfile

	db.Debug().First(&user_profile, 1)
	db.Debug().Model(&user_profile).Association("User").Find(&user_profile.User)
	P("++++++++++++ 1th")
	P(user_profile)

	// 2 Preload Query
	P("++++++++++++ 2th")
	var user_profile2 relate_tables.UserProfile
	db.Preload("User").Find(&user_profile2, 2)
	P(user_profile2)

	// 3 Related Query is invalid
	P("++++++++++++ 3th")
	var user_profile3 relate_tables.UserProfile
	db.First(&user_profile3, 1)
	P(user_profile3)

	var user relate_tables.User
	db.Find(&user, 2)
	P(user)
	// db.Model(&user_profile3).Related(&user, "User")	// err

	// 更新：一定要加条件 // 先关联查询出来，再更新关联表中的字段
	var user_profile4 relate_tables.UserProfile
	db.Preload("User").First(&user_profile4, 2)
	P("++++++++++++ update operate")
	P(user_profile4)

	// single field
	// db.Model(&user_profile4.User).Update("name", "bob1")

	// multiple field
	db.Model(&user_profile4.User).Updates(relate_tables.User{Name: "bob2", Age: 20, Addr: "guanddong2"})

	// Delete Operation
	// 先查询关联， 再删除操作
	// var user_profile5 relate_tables.UserProfile
	// db.Preload("User").Find(&user_profile5, 1)
	// P("++++++++++++ delete operate")
	// P(user_profile5)

	// db.Debug().Delete(&user_profile5.User)
}
