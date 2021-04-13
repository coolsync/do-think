package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var p = fmt.Println

func main() {

	//  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	// db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})

	// First
	// 1
	var user relate_tables.User
	db.First(&user) // 默认使用id查询
	p(user)

	// 2
	var user2 relate_tables.User
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

	// SELECT * FROM `users` WHERE `users`.`name` = 'paul' AND `users`.`age` = 20 AND `users`.`addr` = 'guangdong guangzhou' ORDER BY `users`.`id` LIMIT 1
	tx2 := db.Debug().FirstOrCreate(&user4, user5)
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
	p("user8: ", user8)
}
