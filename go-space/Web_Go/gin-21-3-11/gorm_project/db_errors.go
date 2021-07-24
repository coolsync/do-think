package main

import (
	"errors"
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

	var users []relate_tables.User
	result := db.Where("name", "mark1").Find(&users)

	p(result.RowsAffected)
	p(result.Error)

	// error handle
	if result.Error != nil {
		// do someting
	}

	var user relate_tables.User
	ret := db.Where("name", "mark").First(&user)

	errors.Is(ret.Error, gorm.ErrRecordNotFound) // record not found

	// tx := db.Begin()	// start Transaction
	// ret := db.Commit()

	// if ret.Error != nil {
	// 	tx.Rollback()	// 撤销更改
	// }
	// user2 := relate_tables.User{ID: 11}

	if err := CreateUsers(db); err != nil {
		p("add users failed.")
	}
}

func CreateUsers(db *gorm.DB) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	// add users failed. 已有 id 11， 后面不执行， 直接撤销
	if err := tx.Create(&relate_tables.User{ID: 11, Name: "mark222", Age: 40, Addr: "xx"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Create(&relate_tables.User{ID: 15, Name: "mark222", Age: 40, Addr: "xx"}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
