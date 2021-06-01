package main

import (
	relatetables "comegorm/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	// db.AutoMigrate(&models.User{}, &models.GormModel{}, &models.UserInfo{})

	// one to one belong to
	// db.AutoMigrate(&relatedtables.User1{}, &relatedtables.UserProfile1{})

	// has one, 有外键的先迁移
	// db.AutoMigrate(&relatedtables.UserProfile2{},&relatedtables.User2{})

	// one to many
	// db.AutoMigrate(&relatedtables.UserInfo{}, &relatedtables.CreditCard{})

	// many to many
	db.AutoMigrate(&relatetables.Article{}, &relatetables.Tag{})
}
