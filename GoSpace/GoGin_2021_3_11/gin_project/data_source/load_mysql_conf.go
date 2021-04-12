package data_source

import (
	"encoding/json"
	"fmt"
	"os"
)

// {
//     "host": "localhost",
//     "port": "3306",
//     "user_name": "root",
//     "password": "afvRdOxt%2px",
//     "data_base": "gin_project"
//	   "log_mode": "info"
// }

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
