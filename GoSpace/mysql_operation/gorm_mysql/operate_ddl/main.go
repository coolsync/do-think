package main

import (
	"comegorm/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// Create table for `User`
	db.Migrator().CreateTable(&models.User{})

	// Append "ENGINE=InnoDB" to the creating table SQL for `User`
	// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&models.User{})

	ok := db.Migrator().HasTable(&models.User{})
	// ok := db.Migrator().HasTable("users")
	fmt.Println(ok)

	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable("users")

	// Rename old table to new table
	// db.Migrator().RenameTable(&User{}, &UserInfo{})
	db.Migrator().RenameTable("users", "user_info")
}
