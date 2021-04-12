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

	// INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('paul',20,'xxx',0)
	//  SELECT * FROM `users` WHERE `users`.`name` = 'paul' AND `users`.`age` = 20 AND `users`.`addr` = 'xxx' ORDER BY `users`.`id` LIMIT 1

	p("user4: ", user4)

}
