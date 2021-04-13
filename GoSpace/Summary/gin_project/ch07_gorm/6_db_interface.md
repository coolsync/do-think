# DB接口

# DB 1

## **First**

```go
// 按照主键顺序的第一条记录，（主键升序）

// First
	// 1
	var user relate_tables.User
	db.First(&user) // 按照 default id query
	p(user)

	// 2
	var user2 relate_tables.User
	db.Debug().First(&user2, "name = ?", "bob")	
	// SELECT * FROM `users` WHERE name = 'bob' ORDER BY `users`.`id` LIMIT 1
	p(user2)

	// 3 Where
	var user3 relate_tables.User
	tx1 := db.Debug().Where("name = ?", "bob").First(&user3)
	// SELECT * FROM `users` WHERE name = 'bob' ORDER BY `users`.`id` LIMIT 1
	p(user3)
	p(tx1.RowsAffected) // 返回找到的记录数
	p(tx1.Error)        // Return Err info
	p(user3.ID)         // 返回的 primary key

// sql语句：SELECT * FROM users ORDER BY id LIMIT 1;
```

## **FirstOrCreate**

```go
	// 未找到 user，则根据给定条件创建一条新纪录
	var user4 relate_tables.User

	user5 := relate_tables.User{
		Name: "paul",
		Age:  20,
		Addr: "guangdong guangzhou",
	}

	// SELECT * FROM `users` WHERE `users`.`name` = 'paul' AND `users`.`age` = 20 AND `users`.`addr` = 'guangdong guangzhou' ORDER BY `users`.`id` LIMIT 1
	tx2 := db.Debug().FirstOrCreate(&user4, user5)
	p("user4: ", user4)
	p(tx2.RowsAffected)
```

## **Last**

```go
	// 获取最后一条记录（主键降序）
	var user6 relate_tables.User
	db.Debug().Last(&user6) // SELECT * FROM `users` ORDER BY `users`.`id` DESC LIMIT 1
	p("user6: ", user6)
```

## Take

```go
// 获取一条记录，没有指定排序字段
var user7 relate_tables.User
db.Debug().Take(&user7, 2) // SELECT * FROM `users` WHERE `users`.`id` = 2 LIMIT 1
p("user7: ", user7)
```

## **Find**

```go
// Find
// 多个记录
var user8 []relate_tables.User
id_arr := []int{1, 2, 3}
db.Debug().Find(&user8, id_arr) // SELECT * FROM `users` WHERE `users`.`id` IN (1,2,3)
p("user8: ", user8)


result := db.Find(&users)
// sql语句：SELECT * FROM users;

// 根据指定条件查询
db.Find(&user, "name = ?", "hallen")

//或者结合where
db.Where("name = ?", "hallen").Find(&users)
// sql语句：SELECT * FROM users WHERE name = 'hallen';

db.Where("name LIKE ?", "%ha%").Find(&users)
// sql语句：SELECT * FROM users WHERE name LIKE '%hal%';
```





# DB 2

## **Where**

```go
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
```

## Select

指定要从数据库检索的字段，默认情况下，将选择所有字段;

```go
	var user6 []relate_tables.User
	db.Debug().Select("name, age").Find(&user6) // SELECT name, age FROM `users`
	// db.Select([]string{"name", "age"}).Find(&user6)
	p("user6: ", user6)

	// COALESCE:聚合 ---> ? user7:  {0  0  0}
	var user7 relate_tables.User
	db.Debug().Table("users").Select("COALESCE(age,?)", 30).Rows() //SELECT COALESCE(age,20) FROM `users`
	p("user7: ", user7)
```

```shell
SubQuery子查询

A subquery can be nested within a query, GORM can generate subquery when using a *gorm.DB object as param
子查询可以嵌套在查询中，当使用* gorm.DB对象作为参数时，GORM可以生成子查询

db.Where("amount > (?)", db.Table("orders").Select("AVG(amount)")).Find(&orders)
// SELECT * FROM "orders" WHERE amount > (SELECT AVG(amount) FROM "orders");

subQuery := db.Select("AVG(age)").Where("name LIKE ?", "name%").Table("users")

db.Select("AVG(age) as avgage").Group("name").Having("AVG(age) > (?)", subQuery).Find(&results)
// SELECT AVG(age) as avgage FROM `users` GROUP BY `name` HAVING AVG(age) > (SELECT AVG(age) FROM `users` WHERE name LIKE "name%")
```



## Create

1.插入单条

```go
user := models.User{Name:"李四",Age:18,Addr:"xxx",Pic:"/static/upload/pic111.jpg",Phone:"13411232312"}
result := db.Create(&user)

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
```

2.批量插入：暂不支持

```
user4 := []relate_tables.User{
        {
            Name:"hallen8",
            Age:18,
            Addr:"xxx",
        },
        {
            Name:"hallen9",
            Age:18,
            Addr:"xxx",
        },
    }

db.Create(&user4)    // 这种方式不支持
```

## save

```go
var user model.User

db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
```

## update

```go
var users []model.User
db.Where("active = ?", true).find(&users).Update("name", "hello")


db.Where("active = ?", true).find(&users).Updates(User{Name: "hello", Age: 18})


// update也可以使用map：map[string]interface{}{"name": "hello", "age": 18}

// 也可以使用save更新
```

## delete

```go
db.Delete(&user,1)

// 批量删除
db.Where("age = ?", 20).Delete(&User{})
```

## Unscoped:软删除

```go
    // 也就是逻辑删除    
    // gorm.Model 将DeletedAt 字段设置为当前时间
    // 需要再模型中指定

    type User struct {
      ID      int
      Deleted `gorm:"DeletedAt"`      // 如果设置了所有的删除都将是逻辑删除
      Name    string
    }

    // 在查询时会忽略被软删除的记录
    db.Where("age = 20").Find(&user)
    // SELECT * FROM users WHERE age = 20 AND deleted_at IS NULL;


    // 查询逻辑删除的数据
    db.Unscoped().Where("age = 20").Find(&users)
    // SELECT * FROM users WHERE age = 20;

    // 想要物理删除的办法
    db.Unscoped().Delete(&user)
```

## **Not**

```go
var user model.User

db.Not(User{Name: "hallen", Age: 18}).First(&user)

// SELECT * FROM `users`  WHERE (`users`.`name` <> 'hallen6') AND (`users`.`age` <> 19);
```

## **Or**

```go
var users []model.User


db.Where("name = 'hallen'").Or(User{Name: "hallen2", Age: 18}).Find(&users)
// SELECT * FROM users WHERE name = 'hallen' OR (name = 'jinzhu 2' AND age = 18);
```

## **Order**

```go
var users []model.User
db.Order("age desc").Find(&users) // 注意这里的order要在find前面，否则不生效
fmt.Println(users)

// SELECT * FROM users ORDER BY age desc;

默认为asc
```

## **Limit和Offset**

Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量

```go
var users []model.User


db.Limit(3).Find(&users)  // 三条
// SELECT * FROM users LIMIT 3;

db.Limit(10).Offset(5).Find(&users) // 从5开始的10条数据
// SELECT * FROM users OFFSET 5 LIMIT 10;
```

## **Scan**

将结果扫描到另一个结构中。

```go
type Result struct {
        Id int64
    }
var results []Result
db.Select("id").Where("user_id in (?)", []string{"1", "2"}).Find(&dqmUserRole20).Scan(&results)
fmt.Println(results)
```

## Count

获取模型的记录数

```go
db.Where("name = ?", "hallen").Find(&users).Count(&count)
// SELECT count(*) FROM users WHERE name = 'jinzhu'

db.Model(&User{}).Where("name = ?", "jinzhu").Count(&count)
// SELECT count(*) FROM users WHERE name = 'jinzhu'; (count)

db.Table("deleted_users").Count(&count)
// SELECT count(*) FROM deleted_users;
```

## Group & Having

GROUP BY语句用来与聚合函数(aggregate functions such as COUNT, SUM, AVG, MIN, or MAX.)联合使用，**只返回一个单个值**

HAVING语句通常与GROUP BY语句联合使用，用来过滤由GROUP BY语句返回的记录集。

HAVING语句的存在弥补了WHERE关键字不能与聚合函数联合使用的不足

```go
    type result struct {
      Date  time.Time
      Total int
    }

    db.Select("name, count(*)").Group("name").Find(&result)

    // select name,count(*) FROM users GROUP BY `age`

    db.Select("name, count(*)").Group("name").Having("add = ?","xxx").Find(&result)


  // select name,count(*) FROM users GROUP BY `age` 后面不能用where限制条件，只能使用having

 // select name,count(age) FROM users GROUP BY `age` HAVING addr='xxx'
```

## Distinct：暂无 

```go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

## Join

left join ... on ..

right join ... on ..

```go
db.Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
```

##  

# 