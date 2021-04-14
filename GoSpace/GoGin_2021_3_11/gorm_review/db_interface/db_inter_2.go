package db_interface

import "recome/models/relate_tables"

func DBInter2() {
	// Where
	// 根据条件查询得到满足条件的第一条记录
	var user relate_tables.User
	db.Debug().Where("name", "bob").First(&user) // SELECT * FROM `users` WHERE `name` = 'bob' ORDER BY `users`.`id` LIMIT 1
	p("user: ", user)

	// 根据条件查询得到满足条件的所有记录
	var user2 []relate_tables.User
	db.Debug().Where("name", "bob").Find(&user2) // SELECT * FROM `users` WHERE `name` = 'bob'
	p("user2: ", user2)

	// like模糊查询
	var user3 []relate_tables.User
	db.Debug().Where("name LIKE ?", "p%").Find(&user3) // SELECT * FROM `users` WHERE name LIKE 'p%'
	p("user3: ", user3)

	// 条件
	var user4 []relate_tables.User
	db.Debug().Where("age < ?", 30).Find(&user4) // SELECT * FROM `users` WHERE age < 30
	p("user4: ", user4)

	var user5 []relate_tables.User
	db.Debug().Where("age < ? AND name LIKE ?", 32, "b%").Find(&user5) // SELECT * FROM `users` WHERE age < 32 AND name LIKE 'b%'
	p("user5: ", user5)

	// Select 指定要从数据库检索的字段，默认情况下，将选择所有字段;
	var user6 []relate_tables.User
	db.Debug().Select("name, age").Find(&user6) // SELECT name, age FROM `users`
	// sel := []string{"name", "age"}
	// db.Debug().Select(sel).Find(&user6)
	p("user6: ", user6)

	// db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
	// subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")
	// db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)

	// Get 大于平均年龄的element
	var user7 []relate_tables.User
	db.Debug().Where("age > (?)", db.Table("users").Select("AVG(age)")).Find(&user7) // SELECT * FROM `users` WHERE age > (SELECT AVG(age) FROM `users`)
	p("user7: ", user7)

	// Create	a recode
	var user8 relate_tables.User
	p("user8: ", user8)
}
