package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var p = fmt.Println

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

	// First
	// 1
	var user relate_tables.User
	db.First(&user) // 默认使用id查询
	p(user)

	// 2
	var user2 relate_tables.User
	db.First(&user2, "name = ?", "bob2")
	p(user2)

	// 3
	var user3 relate_tables.User
	ret := db.Where("name = ?", "bob2").First(&user3)
	p(user3)
	p(ret.RowsAffected) // 影响的行数
	p(ret.Error)        // Err Info
	p(user.ID)          // Return Primary Key

	// FirstOrCreate
	// 1

	var user_profile relate_tables.User

	user_profile2 := relate_tables.UserProfile{
		Pic:   "11.jpg",
		CPic:  "22.jpg",
		Phone: "33345678909",
		User: relate_tables.User{
			Name: "bob3",
			Age:  30,
			Addr: "guanddong3",
		},
	}

	ret1 := db.FirstOrCreate(&user_profile, user_profile2)
	p(user_profile)
	p(ret1.RowsAffected)
	// {2 bob2 20 guanddong2 2}

	// Last
	var user4 relate_tables.User
	db.Last(&user4)
	p(user4)

	// Take
	// 获取一条记录，没有指定排序字段
	var user5 relate_tables.User
	db.Take(&user5, 2) // id = 2
	p("user5: ", user5)

	// Find
	// 所有记录
	var user6 relate_tables.User
	id_arr := []int{1, 2, 3, 4}
	db.Find(&user6, id_arr)
	p("user6: ", user6)
	// sql语句：// SELECT * FROM users WHERE id IN (1,2,3, 4);

	// ret3 := db.Find(&user5)
	// sql语句：SELECT * FROM users;

	// 结合where
	var user7 relate_tables.User
	db.Where("name = ?", "bob").Find(&user7)
	p("user7: ", user7)

	var user8 relate_tables.User
	db.Where("name LIKE ?", "%o%").Find(&user8) // only query a recode
	p("user8: ", user8)
}
