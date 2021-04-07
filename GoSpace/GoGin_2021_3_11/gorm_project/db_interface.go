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
<<<<<<< HEAD
=======
	// db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})
>>>>>>> temp

	// First
	// 1
	var user relate_tables.User
<<<<<<< HEAD
	db.First(&user) // 默认使用id查询
=======
	db.First(&user) // 按照 default id query
>>>>>>> temp
	p(user)

	// 2
	var user2 relate_tables.User
<<<<<<< HEAD
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
=======
	db.Debug().First(&user2, "name = ?", "bob")
	// SELECT * FROM `users` WHERE name = 'bob' ORDER BY `users`.`id` LIMIT 1
	p(user2)

	// 3 Where
	var user3 relate_tables.User
	tx1 := db.Debug().Where("name = ?", "bob").First(&user3)
	// SELECT * FROM `users` WHERE name = 'bob' ORDER BY `users`.`id` LIMIT 1
	p(user3)
	p(tx1.RowsAffected) // 返回找到的记录数
	p(tx1.Error)        // Return Err info
	p(user3.ID)         // 返回的 primary key

	// FirstOrCreate
	// 未找到 user，则根据给定条件创建一条新纪录
	var user4 relate_tables.User

	user5 := relate_tables.User{
		Name: "paul",
		Age:  20,
		Addr: "guangdong guangzhou",
	}

	tx2 := db.FirstOrCreate(&user4, user5)
	p("user4: ", user4)
	p(tx2.RowsAffected)

	// Last
	// 获取最后一条记录（主键降序）
	var user6 relate_tables.User
	db.Debug().Last(&user6) // SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	p("user6: ", user6)

	// Take
	// 获取一条记录，没有指定排序字段
	var user7 relate_tables.User
	db.Debug().Take(&user7, 2) // SELECT * FROM `users` WHERE `users`.`id` = 2 LIMIT 1
	p("user7: ", user7)

	// Find
	// 多个记录
	var user8 []relate_tables.User
	id_arr := []int{1, 2, 3}
	db.Debug().Find(&user8, id_arr) // SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
>>>>>>> temp
	p("user8: ", user8)
}
