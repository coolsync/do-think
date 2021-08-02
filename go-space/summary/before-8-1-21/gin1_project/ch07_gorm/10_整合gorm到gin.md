# 配置文件

## 一、配置信息：mysql.json

```json
{
    "host": "localhost",
    "port": "3306",
    "user_name": "root",
    "password": "afvRdOxt%2px",
    "data_base": "gin_project",
    "log_mode": "info"
}
```

## 二、加载配置文件

```go
   package data_source

import (
	"encoding/json"
	"fmt"
	"os"
)

type MysqlConf struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	DataBase string `json:"data_base"`
	LogMode  string `json:"log_mode"`
}

func LoadMysqlConf() *MysqlConf {
	var mysql_conf *MysqlConf

	// get bytes
	bs, err := os.ReadFile("./conf/mysql_conf.json")
	if err != nil {
		panic(err)
	}

	// json unmarshel
	if err := json.Unmarshal(bs, &mysql_conf); err != nil {
		fmt.Println("json unmarshal failed")
		return nil
	}

	return mysql_conf
}
```

## 三、使用配置信息

```go
package data_source

import (
	"fmt"
	"ginproject/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
var err error

func init() {
	mysql_load := LoadMysqlConf()

	level_map := map[string]logger.LogLevel{
		"silent": logger.Silent,	// 	Silent LogLevel = iota + 1
		"error":  logger.Error,
		"warn":   logger.Warn,
		"info":   logger.Info,
	}
	log_mode := mysql_load.LogMode

	// dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gin_project?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysql_load.UserName,
		mysql_load.Password,
		mysql_load.Host,
		mysql_load.Port,
		mysql_load.DataBase,
	)

	// LogMode	Gorm有内置的日志记录器支持 默认情况下，它会打印发生的错误。
	// Logger: logger.Default.LogMode(logger.Info), // global print all sql statements
    Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(level_map[log_mode]), // global print all sql statements
	})
	if err != nil {
		panic(err)
	}
    
	Db.AutoMigrate(&models.User{})

    //Db.DB().SetMaxOpenConns(100) // 最大连接数
    //Db.DB().SetMaxIdleConns(50) // 最大空闲数
}
```