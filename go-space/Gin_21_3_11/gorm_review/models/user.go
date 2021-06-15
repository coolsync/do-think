package models

import "time"

type User struct {
	Id    int
	Name  string
	Age   uint8
	Addr  string
	Pic   string
	Phone string
}

type UserInfo struct {
	Id            int
	Name          string
	DBACreateTime time.Time
}

type DBXXXUserInfo struct { // dbxxx_user_infos
	Id   int
	Name string
}
