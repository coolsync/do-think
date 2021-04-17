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

	// Create

	// Insert a obj
	// user8 := relate_tables.User{
	// 	Name: "sam",
	// 	Age:  30,
	// 	Addr: "xxxx",
	// }
	// db.Debug().Create(&user8) // INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('sam',30,'xxxx',0)

	// 批量插入
	// user9 := []relate_tables.User {
	// 	{
	// 		Name: "jerry4",
	// 		Age: 18,
	// 		Addr: "xxxx",
	// 	},
	// 	{
	// 		Name: "jerry5",
	// 		Age: 20,
	// 		Addr: "xxxx",
	// 	},
	// }
	// db.Debug().Create(&user9)	// INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('jerry4',18,'xxxx',0),('jerry5',20,'xxxx',0)

	// Save
	// user10 := relate_tables.User{
	// 	Name: "sam",
	// 	Age:  32,
	// 	Addr: "xxxx",
	// }
	// db.Debug().Save(&user10) // INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('sam',32,'xxxx',0)

	// var user11 relate_tables.User
	// db.Debug().Where("name", "sam").First(&user11)
	// p(user11)
	// user11.Name = "paul2"
	// db.Debug().Save(&user11) // UPDATE `users` SET `name`='paul2',`age`=30,`addr`='xxxx',`p_id`=0 WHERE `id` = 18

	// Update
	// var user12 relate_tables.User
	// db.Debug().Model(&user12).Where("name", "sam").Update("name", "paul3")	// UPDATE `users` SET `name`='paul3' WHERE `name` = 'sam' // all update
	// p(user12)

	var user13 relate_tables.User
	db.Debug().Where("name", "mark3").Find(&user13).Update("name", "mark4") // UPDATE `users` SET `name`='mark4' WHERE `name` = 'mark3' AND `id` = 22

	db.Debug().Where("name", "mark4").Find(&user13).Updates(relate_tables.User{
		Name: "mark5",
		Age:  30,
	}) // UPDATE `users` SET `name`='mark5',`age`=30 WHERE `name` = 'mark4' AND `id` = 22

	db.Debug().Where("name", "mark5").Find(&user13).Updates(map[string]interface{}{
		"name": "mark6",
		"age":  40,
	}) // UPDATE `users` SET `age`=40,`name`='mark6' WHERE `name` = 'mark5' AND `id` = 22
	p("user13: ", user13)

	// Delete
	var user14 relate_tables.User
	db.Debug().Where("name", "mark6").Delete(&user14) // DELETE FROM `users` WHERE `name` = 'mark6'
	p("user14: ", user14)

	// Not
	var user15 []relate_tables.User
	db.Debug().Not("name", "bob1").Find(&user15) // SELECT * FROM `users` WHERE `name` <> 'bob1'
	p("user15: ", user15)

	// Or
	var user16 []relate_tables.User
	db.Debug().Where("name", "bob1").Or("name", "paul").Find(&user16) // SELECT * FROM `users` WHERE `name` = 'bob1' OR `name` = 'paul'
	p("user16: ", user16)

	// Order
	// desc: from big to small, asc(defualt) : from small to big.
	var user17 []relate_tables.User
	// db.Debug().Order("age desc").Find(&user17) //SELECT * FROM `users` ORDER BY age desc
	db.Debug().Where("name LIKE ?", "b%").Order("id desc").Find(&user17) // SELECT * FROM `users` WHERE name LIKE 'b%' ORDER BY id desc
	p("user17: ", user17)

	// Limit,  Offset(2): from 2+1 start
	var user18 []relate_tables.User
	// db.Debug().Limit(3).Find(&user18)	// SELECT * FROM `users` LIMIT 3
	db.Debug().Limit(5).Offset(2).Find(&user18) // SELECT * FROM `users` LIMIT 5 OFFSET 2,
	p("user18: ", user18)

}
