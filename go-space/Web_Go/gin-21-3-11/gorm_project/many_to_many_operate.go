package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	// 自动移居， 创建两张表
	// db.AutoMigrate(&relate_tables.Article2{}, &relate_tables.Tag{})

	// 1. Insert recodes
	// tag := relate_tables.Tag{
	// 	Name: "标签 1",
	// 	Desc: "标签描述 1",
	// }

	// tag2 := relate_tables.Tag{
	// 	Name: "标签 2",
	// 	Desc: "标签描述 2",
	// }

	// article := relate_tables.Article2{
	// 	Title:   "Article 标题",
	// 	Content: "Article 内容",
	// 	Desc:    "Article 描述",
	// 	Tags: []relate_tables.Tag{
	// 		tag,
	// 		tag2,
	// 	},
	// }

	// 先从db query data, 再 bind 标签
	// var tag3 relate_tables.Tag
	// db.First(&tag3, 3)

	// article := relate_tables.Article2{
	// 	Title:   "Article 标题",
	// 	Content: "Article 内容",
	// 	Desc:    "Article 描述",
	// 	Tags: []relate_tables.Tag{
	// 		tag3,
	// 	},
	// }

	// db.Create(&article)

	// 2. Query data
	var article2 relate_tables.Article2
	db.Preload("Tags").Find(&article2, 1)
	fmt.Println(article2)

	var article3 relate_tables.Article2
	db.First(&article3, 1)
	db.Model(&article3).Association("Tags").Find(&article3.Tags)
	fmt.Println(article3)

	// 3. Update
	// // 先关联查询, 存到 article4
	// var article4 relate_tables.Article2
	// db.Preload("Tags").Find(&article4, 1)

	// // 再根据模型更新， 加条件
	// db.Model(&article4.Tags).Where("id = ?", 1).Update("name", "beego")

	// 4. Delete

	var article5 relate_tables.Article2
	db.Preload("Tags").Find(&article5, 1)
	fmt.Println(article5)
	// db.Where("name = ?", "beego").Delete(&article5.Tags) // invaid

	db.Select(clause.Associations).Where("name = ?", "beego").Delete(&article5.Tags)

	// 解决 Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`gorm_project`.`article2_tags`, CONSTRAINT `fk_article2_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`))
	// 先判断删除关联数据，然后再删除(这样比较符合业务逻辑比较安全)。
}
