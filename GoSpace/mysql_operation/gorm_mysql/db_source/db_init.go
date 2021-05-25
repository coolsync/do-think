package dbsource

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
	err error
)

func init() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	fmt.Println("connect ok")
}
