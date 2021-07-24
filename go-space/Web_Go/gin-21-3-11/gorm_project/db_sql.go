package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var p = fmt.Println

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	//查询一行
	var users []relate_tables.User
	db.Raw("select id, name from users where name = ?", "bob").Find(&users)
	p(users)

	// 增改删用 Exec
	// db.Exec("insert into users (name, age) values(?, ?)", "mark333", 32)
	// db.Exec("update users set name = ? where id = ?", "mark222", 14)
	db.Exec("delete from users where name = ? and age = ?", "mark222", 32)
	

}
