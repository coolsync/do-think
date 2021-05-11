# Many to Many



```go
package relate_tables

// Many To Many
type Article2 struct {
	ID      int
	Title   string
	Content string
	Desc    string
	Tags    []Tag `gorm:"many2many:Article2_tags"` // ;foreignKey:AId;AssociationForeignKey:TId
}

type Tag struct {
	ID   int
	Name string
	Desc string
}

// Many to Many 会在两个 model 中添加一张连接表。

// 例如，您的应用包含了 user 和 language，且一个 user 可以说多种 language，多个 user 也可以说一种 language。

// User 拥有并属于多种 language，`user_languages` 是连接表
// type User3 struct {
// 	gorm.Model
// 	Languages []Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
// 	gorm.Model
// 	Name string
// }

// 反向引用
// type User3 struct {
// 	gorm.Model
// 	Languages []*Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
// 	gorm.Model
// 	Name string
// 	User3s []*User3 `gorm:"many2many:user_languages;"`
// }

// 重写外键
// type User3 struct {
//     gorm.Model
//     Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
//     Refer    uint      `gorm:"index:,unique"`
// }

// type Profile struct {
//     gorm.Model
//     Name      string
//     UserRefer uint `gorm:"index:,unique"`
// }

// 这会创建连接表：user_profiles
//   外键：user_refer_id,，引用：users.refer
//   外键：profile_refer，引用：profiles.user_refer

// CONSTRAINT `fk_user_profiles_profile` FOREIGN KEY (`user_refer`) REFERENCES `profiles` (`user_refer`),
// CONSTRAINT `fk_user_profiles_user3` FOREIGN KEY (`user_refer_id`) REFERENCES `user3` (`refer`)

```



# Many to Many operate



```go
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
	
}
```



解决 1: 

```shell
Error 1451: Cannot delete or update a parent row: a foreign key constraint fails (`gorm_project`.`article2_tags`, CONSTRAINT `fk_article2_tags_tag` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`))

先判断删除关联数据，然后再删除(这样比较符合业务逻辑比较安全)。
```



解决 2: 

------

要全局关闭外键约束，请执行以下操作：

| 1    | [SET](http://search.oracle.com/search/search?group=MySQL&q=SET) [GLOBAL](http://search.oracle.com/search/search?group=MySQL&q=GLOBAL) FOREIGN_KEY_CHECKS=0; |
| ---- | ------------------------------------------------------------ |
|      |                                                              |

并记得在完成后将其设置回来

| 1    | [SET](http://search.oracle.com/search/search?group=MySQL&q=SET) [GLOBAL](http://search.oracle.com/search/search?group=MySQL&q=GLOBAL) FOREIGN_KEY_CHECKS=1; |
| ---- | ------------------------------------------------------------ |
|      |                                                              |

警告：只有在进行单用户模式维护时才应执行此操作。因为它可能导致数据不一致。例如，当您使用mysqldump输出上载大量数据时，它将非常有用。



