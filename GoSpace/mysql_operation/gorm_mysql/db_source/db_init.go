package dbsource

import (
	"comegorm/models"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
	p   = fmt.Println
)

func init() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	Db.AutoMigrate(&models.User{})

	sqlDB, err := Db.DB()
	if err != nil {
		log.Fatal("Db.DB() err: ", err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	// 返回数据库统计信息
	// sqlDB.Stats()

	p("connect ok")
}
