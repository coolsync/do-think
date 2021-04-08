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

	// Where
	// 根据条件查询得到满足条件的第一条记录
	var user []relate_tables.User
	db.Debug().Where("name", "bob").First(&user)
	// SELECT * FROM `users` WHERE `name` = 'bob' ORDER BY `users`.`id` LIMIT 1
	p(user)

	// 根据条件查询得到满足条件的所有记录
	var user2 []relate_tables.User
	db.Debug().Where("name", "bob").Find(&user2) //  SELECT * FROM `users` WHERE `name` = 'bob'
	p(user2)

	// like模糊查询
	var user3 []relate_tables.User
	db.Debug().Where("name like ?", "p%").Find(&user3) // SELECT * FROM `users` WHERE name like 'p%'
	p(user3)

	var user4 []relate_tables.User
	db.Debug().Where("age < ?", 30).Find(&user4) //  SELECT * FROM `users` WHERE age < 30
	p(user4)

	// 	条件：
	var user5 []relate_tables.User
	db.Debug().Where("name = ? AND age >= ?", "bob2", 20).Find(&user5)
	// SELECT * FROM `users` WHERE name = 'bob2' AND age >= 20
	p(user5)

	// =
	// LIKE
	// IN：Where("name IN ?", []string{"bob2", "paul"})
	// AND：Where("name = ? AND age >= ?", "jinzhu", "22")
	// Time：Where("updated_at > ?", lastWeek)
	// BETWEEN：Where("created_at BETWEEN ? AND ?", lastWeek, today)

	p("+++++++++++++++++++++++ Select")
	// Select
	// 指定要从数据库检索的字段，默认情况下，将选择所有字段;

	var user6 []relate_tables.User
	db.Debug().Select("name, age").Find(&user6) // SELECT name, age FROM `users`
	// db.Select([]string{"name", "age"}).Find(&user6)
	p("user6: ", user6)

	// COALESCE:聚合 ---> ? user7:  {0  0  0}
	var user7 relate_tables.User
	db.Debug().Table("users").Select("COALESCE(age,?)", 30).Rows() //SELECT COALESCE(age,20) FROM `users`
	p("user7: ", user7)

	// Create
	// 1.插入单条
	// user8 := relate_tables.User{
	// 	Name:"jerry",
	// 	Age: 18,
	// 	Addr: "xxxx",
	// }

	// db.Create(&user8)

	// 2.批量插入
	// user9 := []relate_tables.User {
	// 	{
	// 		Name:"jerry2",
	// 		Age: 18,
	// 		Addr: "xxxx",
	// 	},
	// 	{
	// 		Name:"jerry3",
	// 		Age: 18,
	// 		Addr: "xxxx",
	// 	},
	// }

	// db.Create(&user9)

	// Save
	// user10 := relate_tables.User{
	// 	Name: "mark",
	// 	Age: 30,
	// 	Addr: "xxx",
	// }
	// db.Save(&user10)

	// var user11 relate_tables.User
	// db.Where("name", "mark").First(&user11)
	// p(user11)

	// user11.Name = "mark1"
	// db.Save(&user11)

	// update

	// Distinct: 除重
	// Selecting distinct values from the model
	// SELECT DISTINCT name, age FROM `users` ORDER BY name, age desc

	var user22 []relate_tables.User
	db.Debug().Distinct("name, age").Order("name, age desc").Find(&user22)
	p("user22: ", user22)

	// Joins
	// select * from users right join user_profiles on users.p_id = user_profiles.id;
	// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`users`.`addr`,`users`.`p_id` FROM `users` right join user_profiles on users.p_id = user_profiles.id
	type UserJoins struct {
		ID    int
		Name  string
		Age   int
		Addr  string
		PID   int
		Pic   string
		CPic  string
		Phone string
	}
	var user_joins []UserJoins
	// var user23 []relate_tables.User

	// db.Debug().Model(&user23).Select("users.*, user_profiles.*").Joins("right join user_profiles on users.p_id = user_profiles.id").Scan(&user_joins)

	db.Debug().Table("users").Select("users.id, users.p_id, user_profiles.pic").Joins("right join user_profiles on users.p_id = user_profiles.id").Scan(&user_joins)

	fmt.Printf("%#v", user_joins)
	// p(user_joins)

}
