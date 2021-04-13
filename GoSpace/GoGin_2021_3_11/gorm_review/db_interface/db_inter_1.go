package db_interface

import (
	"fmt"
	"recome/db_source"
	"recome/models/relate_tables"
)

var p = fmt.Println
var db = db_source.Db

func DBInter1() {
	db.AutoMigrate(&relate_tables.UserProfile{}, &relate_tables.User{})

	// First
	var user relate_tables.User
	db.Debug().First(&user) // SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	p("user: ", user)

	var user2 relate_tables.User
	db.Debug().First(&user2, "name", "bob") // SELECT * FROM `users` WHERE `name` = 'bob' ORDER BY `users`.`id` LIMIT 1
	p("user2: ", user2)

	var user3 relate_tables.User
	db.Debug().Where("name", "jerry").First(&user3) // SELECT * FROM `users` WHERE `name` = 'jerry' ORDER BY `users`.`id` LIMIT 1
	p("user3: ", user3)

	// FirstOrCreate
	var user4 relate_tables.User

	user5 := relate_tables.User{
		Name: "paul",
		Age:  20,
		Addr: "xxx",
	}
	db.Debug().FirstOrCreate(&user4, user5)

	// INSERT INTO `users` (`name`,`age`,`addr`,`p_id`)
	// VALUES ('paul',20,'xxx',0)

	// SELECT * FROM `users`
	// WHERE `users`.`name` = 'paul' AND `users`.`age` = 20 AND `users`.`addr` = 'xxx'
	// ORDER BY `users`.`id` LIMIT 1

	p("user4: ", user4)

	// Last
	var user6 relate_tables.User
	db.Debug().Last(&user6) // SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	p("user6: ", user6)

	// Take
	var user7 relate_tables.User
	db.Debug().Take(&user7, "name", "mark1") // SELECT * FROM `users` WHERE name = 'mark1' LIMIT
	p("user7: ", user7)

	// Find 收集多个recode
	var user8 []relate_tables.User
	id_arr := []int{1, 2, 3}
	db.Debug().Find(&user8, id_arr) // SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
	p("user8: ", user8)

	var user9 []relate_tables.User
	db.Debug().Find(&user9, "name = ? AND age = ?", "bob", 30) // SELECT * FROM `users` WHERE name = 'bob' AND age = 30
	p("user9: ", user9)

	var user10 []relate_tables.User
	db.Debug().Where("age", 18).Find(&user10) // SELECT * FROM `users` WHERE `age` = 18
	p("user10: ", user10)

	var user11 []relate_tables.User
	db.Debug().Where("name LIKE ?", "%y%").Find(&user11) // SELECT * FROM `users` WHERE name LIKE '%y%'
	p("user11: ", user11)
}
