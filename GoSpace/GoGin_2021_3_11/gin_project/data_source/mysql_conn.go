package data_source

import (
	"fmt"
	"ginproject/models"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func init() {
	mysql_load := LoadMysqlConf()

	log_mode, _ := strconv.Atoi(mysql_load.LogMode)

	// dsn1 := "root:afvRdOxt%2px@tcp(localhost:3306)/gin_project?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql_load.UserName,
		mysql_load.Password,
		mysql_load.Host,
		mysql_load.Port,
		mysql_load.DataBase,
	)

	// fmt.Println("========================", dsn)

	// LogMode	Gorm有内置的日志记录器支持 默认情况下，它会打印发生的错误。
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info), // global print all sql statements
		Logger: logger.Default.LogMode(logger.LogLevel(log_mode)), // global print all sql statements
	})

	if err != nil {
		panic(err)
	}

	Db.AutoMigrate(&models.User{})

}
