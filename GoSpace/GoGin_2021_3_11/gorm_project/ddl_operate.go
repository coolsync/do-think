package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

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

	// Create Table
	// db.Migrator().CreateTable(&User{})	// 使用模型名
	// db.Table("user").Migrator().CreateTable(&User{})

	// Delete Table
	// db.Migrator().DropTable(&User{})
	// db.Migrator().DropTable("user")

	// Has Table?
	// b := db.Migrator().HasTable("users")
	// b := db.Migrator().HasTable(&models.User{})
	// fmt.Println(b)

	// b2 := db.Migrator().HasTable("user")
	// fmt.Println(b2)

	// 统一加prefix, suffix

	// 自动迁移
	// db.AutoMigrate(&models.User{}, &models.UserInfo{}, &models.DBXXXUserInfo{})
	// db.AutoMigrate(&models.User{}, &models.GormModel{})

	// db.AutoMigrate(&models.User{}, &models.Article{})
	// db.AutoMigrate(&relate_tables.User{}, &relate_tables.Company{})

	// one to one
	// db.AutoMigrate(&relate_tables.User{}, &relate_tables.UserProfile{})

	// one to many
	// db.AutoMigrate(&relate_tables.User2{}, &relate_tables.Article{})

	// many to many
	// db.AutoMigrate(&relate_tables.User3{}, &relate_tables.Language{})
	// db.AutoMigrate(&relate_tables.User3{}, &relate_tables.Profile{})
	db.AutoMigrate(&relate_tables.Article2{}, &relate_tables.Tag{})

	fmt.Println("DDL OK")
}

