# One to Many



```go
package relate_tables

// one to many
type User2 struct {
	ID int
	Name string
	Age int
	Addr string
	Articles []Article `gorm:"foreignKey:UID;references:ID"`
}

type Article struct {
	ID int
	Title string
	Content string
	Desc string
	UID int // foreign key
}
```





# 一对多操作、

## 一、增加

```go
	// 第一种方式
	user2 := relate_tables.User2{
		Name: "paul",
		Age:  20,
		Addr: "xxxx",
		Articles: []relate_tables.Article{
			{
				Title:   "beego xiangjian 1",
				Content: "beego xiangjian content 1",
				Desc:    "beego xiangjian 1 description",
			},
			{
				Title:   "beego xiangjian 2",
				Content: "beego xiangjian content 2",
				Desc:    "beego xiangjian description 2",
			},
		},
	}
	db.Create(&user2)


	ret := db.Create(&user2)

	fmt.Println(ret.RowsAffected)
	fmt.Println(ret.Error)
```

## 二、查询

1.Preload

```go
var user2 relate_tables.User2

db.Preload("Articles").Find(&user2, 1)
fmt.Println(user2)
```

2.Association

```go
var user3 relate_tables.User2

db.First(&user3, 2)
db.Model(&user3).Association("Articles").Find(&user3.Articles)
fmt.Println(user3)
```

3.Related -> Invaild

```go
var user2 relate_tables.User2

db.First(&user2,1)

var articles []relate_table.Article

db.Model(&user2).Related(&articles, "Articles")     // 关系名称
```

## 三、更新

```go
// 先查询
var user2 relate_tables.User2

db.Preload("Articles").Find(&user2,2)    // 关系名

// 再更新，更新指定条件，不然会把所有满足条件的都更新

db.Model(&user2.Articles).Where("title=? and uid=?","标题测试2",2).Update("uid",3) // name和uid限制条件

// update只能更新一个字段，如果想同时更新多个字段，使用save，后面会讲
```

## 四、删除

```go
// 先查询
var user2 relate_tables.User2
db.Preload("Articles").Find(&user2,2)    // 关系名



// 再删除，删除要指定条件，不然会把所有满足条件的都删除
db.Delete(&user2.Articles,"title=? and uid=?","标题测试3",3)	

// 或者使用where
db.Where("title=? and uid=?","标题测试3",3).Delete(&user2.Articles)
```