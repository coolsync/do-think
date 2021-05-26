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

	// MySQl 驱动程序提供了 一些高级配置 可以在初始化过程中使用，例如：
	// db, err := gorm.Open(mysql.New(mysql.Config{
	// 	DSN:                       "gorm:gorm@tcp(127.0.0.1:3306)/gorm?charset=utf8&parseTime=True&loc=Local", // DSN data source name
	// 	DefaultStringSize:         256,                                                                        // string 类型字段的默认长度
	// 	DisableDatetimePrecision:  true,                                                                       // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
	// 	DontSupportRenameIndex:    true,                                                                       // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
	// 	DontSupportRenameColumn:   true,                                                                       // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
	// 	SkipInitializeWithVersion: false,                                                                      // 根据当前 MySQL 版本自动配置
	// }), &gorm.Config{})

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
