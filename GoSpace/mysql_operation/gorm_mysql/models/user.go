package models

type User struct {
	Id    int
	Name  string
	Age   uint8
	Addr  string
	Pic   string
	Phone string
}

type UserInfo struct {
	Id   int    `gorm:"primary_key"`
	Name string `gorm:"index"`
	Age  int
}

// type DBXXXUserInfo struct { // dbxxx_user_infos
// 	Id   int
// 	Name string
// }
