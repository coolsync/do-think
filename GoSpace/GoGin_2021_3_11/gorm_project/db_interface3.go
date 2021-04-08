package main

import (
	"fmt"
	"gorm_project/models/relate_tables"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var p = fmt.Println

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	// LogMode	Gorm有内置的日志记录器支持 默认情况下，它会打印发生的错误。
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // global print all sql statements
	})

	if err != nil {
		panic(err)
	}

	// 高级查询
	// FirstOrInit： 不会影响 DB data
	// 获取第一个匹配的记录，或者使用给定的条件初始化一个新的记录（仅适用于struct，map条件）
	var user relate_tables.User
	// db.Debug().FirstOrInit(&user)	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	db.FirstOrInit(&user, relate_tables.User{Name: "jerry", Age: 30}) // {0 jerry 30  0}
	p(user)

	// Attrs：如果没有找到记录，则使用Attrs中的数据来初始化一条记录：
	var user3 relate_tables.User
	db.Where("name", "mark").Attrs(relate_tables.User{Name: "mark2", Age: 32}).FirstOrInit(&user3)
	p(user3)

	// Assign:不管是否找的到，最终返回结构中都将带上Assign指定的参数，有则代替，没有则添加
	var user4 relate_tables.User
	db.Where("name", "mark1").Assign(relate_tables.User{Name: "mark1", Age: 32}).FirstOrInit(&user4)
	p(user4)

	// Pluck 用于从数据库查询单个列，并将结果扫描到切片。如果您想要查询多列，您应该使用 Select 和 Scan
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

	// Scopes 允许你指定常用的查询，可以在调用方法时引用这些查询
	// 无 params
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
