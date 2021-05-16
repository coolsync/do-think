package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&relate_tables.User2{}, &relate_tables.Article{})

	// 添加
	// 第一种方式
	// user2 := relate_tables.User2{
	// 	Name: "paul",
	// 	Age:  20,
	// 	Addr: "xxxx",
	// 	Articles: []relate_tables.Article{
	// 		{
	// 			Title:   "beego xiangjian 1",
	// 			Content: "beego xiangjian content 1",
	// 			Desc:    "beego xiangjian 1 description",
	// 		},
	// 		{
	// 			Title:   "beego xiangjian 2",
	// 			Content: "beego xiangjian content 2",
	// 			Desc:    "beego xiangjian description 2",
	// 		},
	// 	},
	// }
	// db.Create(&user2)

	// 第2种方式 推荐使用第一种方式
	// article := relate_tables.Article{
	// 	Title:   "beego xiangjian 3",
	// 	Content: "beego xiangjian content 3",
	// 	Desc:    "beego xiangjian description 3",
	// }
	// // db.Create(&article)

	// article2 := relate_tables.Article{
	// 	Title:   "beego xiangjian 4",
	// 	Content: "beego xiangjian content 4",
	// 	Desc:    "beego xiangjian description 4",
	// }
	// // db.Create(&article2)

	// user2 := relate_tables.User2{
	// 	Name: "paul",
	// 	Age:  20,
	// 	Addr: "xxxx",
	// 	Articles: []relate_tables.Article{
	// 		article,
	// 		article2,
	// 	},
	// }
	// db.Create(&user2)

	// 二、查询
	// Perload
	var user2 relate_tables.User2

	db.Preload("Articles").Find(&user2, 1)
	fmt.Println(user2)

	// Association
	var user3 relate_tables.User2

	db.First(&user3, 2)
	db.Model(&user3).Association("Articles").Find(&user3.Articles)
	fmt.Println(user3)

	// update operate
	fmt.Println("+++++++++++update operate")
	var user4 relate_tables.User2
	db.Preload("Articles").Find(&user4, 1)
	fmt.Println(user4)

	db.Model(&user4.Articles).Where("id=?", 1).Update("title", "beego xiangjian 1")

	// Delete Operate
	// 先查询
	var user5 relate_tables.User2
	db.Preload("Articles").Find(&user5, 1)
	fmt.Println(user5)

	fmt.Println("+++++++++++update operate")

	db.Delete(&user5.Articles, "title=? and uid=?", "beego xiangjian 3", 2)

	// use where
	db.Where("title=? and uid=?", "beego xiangjian 1", 1).Delete(&user5.Articles)
}
