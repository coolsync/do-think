package db_source

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
	err error
)

func init() {
	//  // 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	// dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	// Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_project?charset=utf8mb4&parseTime=True&loc=Local"

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}
}
