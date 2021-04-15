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
//Create
//1.插入单条
user8 := relate_tables.User{
    Name:"jerry",
    Age: 18,
    Addr: "xxxx",
}

db.Create(&user8) // INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('sam',30,'xxxx',0)

user.ID             // 返回插入数据的主键
result.Error        // 返回 error
result.RowsAffected // 返回插入记录的条数
```

2.批量插入

```go
// 2.批量插入
user9 := []relate_tables.User {
    {
        Name: "jerry4",
        Age: 18,
        Addr: "xxxx",
    },
    {
        Name: "jerry5",
        Age: 20,
        Addr: "xxxx",
    },
}
db.Debug().Create(&user9)	// INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('jerry4',18,'xxxx',0),('jerry5',20,'xxxx',0)
```

## Save 

Save update value in database, if the value doesn't have primary key, will insert it

默认会更新该对象的所有字段，即使你没有赋值。

```go
user10 := relate_tables.User{
    Name: "mark",
    Age: 30,
    Addr: "xxx",
}
db.Save(&user10) // INSERT INTO `users` (`name`,`age`,`addr`,`p_id`) VALUES ('sam',32,'xxxx',0)

var user11 relate_tables.User
db.Where("name", "mark").First(&user11) // UPDATE `users` SET `name`='paul2',`age`=30,`addr`='xxxx',`p_id`=0 WHERE `id` = 18
p(user11)
user11.Name = "mark1"
db.Save(&user11)
```



```go
var user model.User

db.First(&user)

user.Name = "jinzhu 2"
user.Age = 100
db.Save(&user)
```

## update

```go
// all update
var user12 relate_tables.User
db.Debug().Model(&user12).Where("name", "sam").Update("name", "paul3")	// UPDATE `users` SET `name`='paul3' WHERE `name` = 'sam' 
p(user12)

// 先查询 后更新
var user12 relate_tables.User
db.Where("name = ?", "mark1").First(&user12)
p(user12)
db.Model(&user12).Update("name", "mark2")

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


// update也可以使用map：map[string]interface{}{"name": "hello", "age": 18}

// 也可以使用save更新
```

## delete

```go
var user14 relate_tables.User
db.Where("name", "jerry3").Delete(&user14)
p("user14:", user14)	// user14: {0  0  0}

// 批量删除
db.Where("email LIKE ?", "%jinzhu%").Delete(Email{})
//// DELETE from emails where email LIKE "%jinhu%";

db.Delete(Email{}, "email LIKE ?", "%jinzhu%")
//// DELETE from emails where email LIKE "%jinhu%";
```



## Unscoped:软删除

```go
    // 也就是逻辑删除    
    // gorm.Model 将DeletedAt 字段设置为当前时间
    // 需要在模型中指定

    type User struct {
      ID      int
      Deleted `gorm:"DeletedAt"`      // 如果设置了 所有的删除 都将是逻辑删除
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
var user15 []relate_tables.User
db.Debug().Not("name", "bob").Find(&user15) // SELECT * FROM `users` WHERE `name` <> 'bob'
p(user15)

var user model.User
db.Not(User{Name: "hallen", Age: 18}).First(&user) // SELECT * FROM `users`  WHERE (`users`.`name` <> 'hallen6') AND (`users`.`age` <> 19);
```

## **Or**

```go
var user16 []relate_tables.User
db.Debug().Where("name", "bob").Or("name", "paul").Find(&user16) //SELECT * FROM `users` WHERE `name` = 'bob' OR `name` = 'paul'
p(user16)

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
// 默认为asc

var user17 []relate_tables.User
db.Debug().Where("name LIKE ?", "b%").Order("name desc").Find(&user17) // SELECT * FROM `users` WHERE name LIKE 'b%' ORDER BY name asc
p(user17)
```

## **Limit和Offset**

Limit 指定获取记录的最大数量 Offset 指定在开始返回记录之前要跳过的记录数量

```go
var users []model.User
db.Limit(3).Find(&users)  // 三条 // SELECT * FROM users LIMIT 3;
db.Limit(10).Offset(5).Find(&users) // 从5开始的10条数据 // SELECT * FROM users OFFSET 5 LIMIT 10;

var user18 []relate_tables.User
// db.Debug().Limit(3).Find(&user18) //SELECT * FROM `users` LIMIT 3
db.Debug().Limit(5).Offset(3).Find(&user18) // SELECT * FROM `users` LIMIT 5 OFFSET 3
p(user18)
```

## **Scan**

将结果扫描到另一个结构中。

```go
type Result struct {
        Id int64
    }
var results []Result
db.Select("id").Where("user_id in (?)",[]string{"1","2"}).Find(&dqmUserRole20).Scan(&results)
fmt.Println(results)

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
```

## Count

获取模型的记录数

```go
var user20 []relate_tables.User
var count int64
// db.Debug().Where("age", 30).Find(&user20).Count(&count) // SELECT count(1) FROM `users` WHERE `age` = 30
db.Debug().Model(&user20).Where("age", 30).Count(&count)
// p(user20)
p(count)


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
type GroupData struct {
    Name  string
    Age   string
    Addr  string
    Count int
}
var group_data []GroupData
var user21 []relate_tables.User
// db.Debug().Model(&user21).Select("age, count(*) as count").Group("age").Find(&group_data)
db.Model(&user21).Select("age, count(*) as count").Group("age").Having("age = ?", 30).Find(&group_data)   


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
var user22 []relate_tables.User
db.Debug().Distinct("name, age").Order("name, age desc").Find(&user22)
p("user22: ", user22)
```

## Join

left join ... on ..

right join ... on ..

```go
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
```

##  
