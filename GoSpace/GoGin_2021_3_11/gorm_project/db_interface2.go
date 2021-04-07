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
	// var user12 relate_tables.User
	// db.Where("name = ?", "mark1").First(&user12)
	// p(user12)
	// db.Model(&user12).Update("name", "mark2")

	// var user13 relate_tables.User
	// db.Where("name", "mark2").Find(&user13).Update("name", "mark3")
	// db.Where("name", "mark3").Find(&user13).Updates(relate_tables.User{Name: "mark4", Age: 40})
	// db.Where("name", "mark4").Find(&user13).Updates(map[string]interface{}{
	// 	"name": "mark5",
	// 	"age": "30",
	// })

	// Delete
	// var user14 relate_tables.User
	// db.Where("name", "jerry3").Delete(&user14)
	// p("user14:", user14)	// user14: {0  0  0}

	// Not
	var user15 []relate_tables.User
	db.Debug().Not("name", "bob").Find(&user15) // SELECT * FROM `users` WHERE `name` <> 'bob'
	p(user15)

	// Or
	var user16 []relate_tables.User
	//SELECT * FROM `users` WHERE `name` = 'bob' OR `name` = 'paul'
	db.Debug().Where("name", "bob").Or("name", "paul").Find(&user16)
	p(user16)

	// Order
	var user17 []relate_tables.User
	// SELECT * FROM `users` WHERE name LIKE 'b%' ORDER BY name asc
	db.Debug().Where("name LIKE ?", "b%").Order("name desc").Find(&user17)
	p(user17)

	// Limit和Offset
	// Limit 指定获取记录的最大数量 Offset(3) 指定在开始返回记录之前要跳过的记录数量 1..3, start 4
	var user18 []relate_tables.User
	// db.Debug().Limit(3).Find(&user18) //SELECT * FROM `users` LIMIT 3
	db.Debug().Limit(5).Offset(3).Find(&user18) // SELECT * FROM `users` LIMIT 5 OFFSET 3
	p(user18)

	// Scan
	// 将结果扫描到另一个结构中。
	type UserBak struct {
		Name string
		// Age int
		Addr string
	}
	var user_bak []UserBak
	var user19 []relate_tables.User
	db.Find(&user19).Scan(&user_bak)
	p(user19)
	p(user_bak)

	// Count
	// 获取模型的记录数
	var user20 []relate_tables.User
	var count int64
	// db.Debug().Where("age", 30).Find(&user20).Count(&count)
	// SELECT count(1) FROM `users` WHERE `age` = 30
	db.Debug().Model(&user20).Where("age", 30).Count(&count)
	// p(user20)
	p(count)

	// Group & Having
	// GROUP BY语句用来与聚合函数(aggregate functions such as COUNT, SUM, AVG, MIN, or MAX.)联合使用，只返回一个单个值
	// HAVING语句通常与GROUP BY语句联合使用，用来过滤由GROUP BY语句返回的记录集。
	// HAVING语句的存在弥补了WHERE关键字不能与聚合函数联合使用的不足

	type GroupData struct {
		Name  string
		Age   string
		Addr  string
		Count int
	}
	var group_data []GroupData
	var user21 []relate_tables.User
	// db.Debug().Model(&user21).Select("age, count(*) as count").Group("age").Find(&group_data)
	// db.Debug().Model(&user21).Select("age, addr, count(*) as count").Group("age").Having("addr = ?", "guangdong shengzheng").Find(&group_data)

	db.Model(&user21).Select("age, count(*) as count").Group("age").Having("age = ?", 30).Find(&group_data)

	// p(user21)
	p(group_data)
}
