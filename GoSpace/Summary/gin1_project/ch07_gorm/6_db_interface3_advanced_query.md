## Distinct

Selecting distinct values from the model

```go
db.Distinct("name", "age").Order("name, age desc").Find(&results)
```

`Distinct` works with [`Pluck`](https://gorm.io/docs/advanced_query.html#pluck), [`Count`](https://gorm.io/docs/advanced_query.html#count) also



## Joins

Specify Joins conditions

left join ... on ..

right join ... on ..

```go
type result struct {
  Name  string
  Email string
}
db.Model(&User{}).Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&result{})
// SELECT users.name, emails.email FROM `users` left join emails on emails.user_id = users.id

rows, err := db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Rows()
for rows.Next() {
  ...
}

db.Table("users").Select("users.name, emails.email").Joins("left join emails on emails.user_id = users.id").Scan(&results)

// multiple joins with parameter
db.Joins("JOIN emails ON emails.user_id = users.id AND emails.email = ?", "jinzhu@example.org").Joins("JOIN credit_cards ON credit_cards.user_id = users.id").Where("credit_cards.number = ?", "411111111111").Find(&user)
```

### Joins Preloading

You can use `Joins` eager loading associations with a single SQL, for example:

```go
db.Joins("Company").Find(&users)
// SELECT `users`.`id`,`users`.`name`,`users`.`age`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` AS `Company` ON `users`.`company_id` = `Company`.`id`;
```

Refer [Preloading (Eager Loading)](https://gorm.io/docs/preload.html) for details





# 高级查询

## **FirstOrInit和Attrs**

**FirstOrInit：获取第一个匹配的记录，或者使用给定的条件初始化一个新的记录（仅适用于struct，map条件）**

```go
var user relate_tables.User
// db.Debug().FirstOrInit(&user)	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

db.FirstOrInit(&user, relate_tables.User{Name: "jerry", Age: 30}) // {0 jerry 30  0}
p(user)
```

Attrs：如果没有找到记录，则使用Attrs中的数据来初始化一条记录：

```go
var user3 relate_tables.User
// 查不到该条记录，则使用attrs值替换,// 查到记录，则使用数据库中的值
db.Where("name", "mark").Attrs(relate_tables.User{Name: "mark2", Age: 32}).FirstOrInit(&user3)
p(user3)
```

## **FirstOrInit和Assign**

Assign:不管是否找的到，最终返回结构中都将带上Assign指定的参数，有则代替，没有则添加

```go
	
var user4 relate_tables.User
// 不管是否找到对应记录，使用Assign值替代查询到的值

db.Where("name", "mark1").Assign(relate_tables.User{Name: "mark1", Age: 32}).FirstOrInit(&user4)
	
p(user4)
```

## Pluck

Query single column from database and scan into a slice, if you want to query multiple columns, use `Select` with [`Scan`](https://gorm.io/docs/query.html#scan) instead

Pluck 用于从数据库查询单个列，并将结果扫描到切片。如果您想要查询多列，您应该使用 Select 和 Scan

```go
var ages []int
	
var users []relate_tables.User
	
db.Model(&users).Where("age > ?", 20).Pluck("age", &ages)
	
p(ages)

var names []string
	
db.Model(&relate_tables.User{}).Pluck("name", &names)

p(names)
	
db.Table("users").Pluck("name", &names)
	
p(names)
		
db.Model(&relate_tables.User{}).Distinct().Pluck("name", &names)
	
p(names)
```



```go
var ages []int64
db.Model(&users).Pluck("age", &ages)

var names []string
db.Model(&User{}).Pluck("name", &names)

db.Table("deleted_users").Pluck("name", &names)

// Distinct Pluck
db.Model(&User{}).Distinct().Pluck("Name", &names)
// SELECT DISTINCT `name` FROM `users`

// Requesting more than one column, use `Scan` or `Find` like this:
db.Select("name", "age").Scan(&users)
db.Select("name", "age").Find(&users)
```



## Scopes

Scopes 允许你指定常用的查询，可以在调用方法时引用这些查询



```go
// 无 params
{	
    var users_on []relate_tables.User
    var users_off []relate_tables.User

    db.Scopes(GetStatusOn).Find(&users_on)
    db.Scopes(GetStatusOff).Find(&users_off)

    p(users_on)
    p(users_off)

    // has params
    var users_names []relate_tables.User
    db.Scopes(GetRowsByName("mark1")).Find(&users_names)
    p(users_names)
}

func GetStatusOn(db *gorm.DB) *gorm.DB {
	//return db.Where("status = ?",1)   // 正常的
	return db.Where("p_id", 1) // 为了演示
}

func GetStatusOff(db *gorm.DB) *gorm.DB {
	//return db.Where("status = ?",1)   // 正常的
	return db.Where("p_id", 0) // 为了演示
}

func GetRowsByName(name string) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("name", name)
	}
}
```



```go
func AmountGreaterThan1000(db *gorm.DB) *gorm.DB {
  return db.Where("amount > ?", 1000)
}

func PaidWithCreditCard(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func PaidWithCod(db *gorm.DB) *gorm.DB {
  return db.Where("pay_mode_sign = ?", "C")
}

func OrderStatus(status []string) func (db *gorm.DB) *gorm.DB {
  return func (db *gorm.DB) *gorm.DB {
    return db.Where("status IN (?)", status)
  }
}

db.Scopes(AmountGreaterThan1000, PaidWithCreditCard).Find(&orders)
// 查找所有金额大于 1000 的信用卡订单

db.Scopes(AmountGreaterThan1000, PaidWithCod).Find(&orders)
// 查找所有金额大于 1000 的货到付款订单

db.Scopes(AmountGreaterThan1000, OrderStatus([]string{"paid", "shipped"})).Find(&orders)
// 查找所有金额大于 1000 且已付款或已发货的订单
```

## LogMode

Gorm有内置的日志记录器支持，默认情况下，它会打印发生的错误。

启用Logger，显示详细日志

```go
db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
    Logger: logger.Default.LogMode(logger.Info),	// global print all sql statements
})
```

查看单语句的信息

```go
db.Debug().Where("name = ?", "hallen").First(&User{})
```