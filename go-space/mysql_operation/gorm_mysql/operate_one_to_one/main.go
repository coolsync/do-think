package main

import (
	relatetables "comegorm/models/relate_tables"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:afvRdOxt%2px@tcp(localhost:3306)/gorm_mysql?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(relatetables.UserProfile2{}, relatetables.User2{})

	// db.Migrator().DropTable(relatetables.User2{})
	// db.Migrator().DropTable(relatetables.UserProfile2{})

	// // Add
	// user_profile := relatetables.UserProfile2{
	// 	Pic:   "2.jpg",
	// 	CPic:  "222.jpg",
	// 	Phone: "12345678",
	// 	User: relatetables.User2{
	// 		Name: "mark",
	// 		Age:  30,
	// 		Addr: "xxxx",
	// 	},
	// }
	// db.Create(&user_profile)

	// Query 1: Association
	var u_profile relatetables.UserProfile2
	db.Debug().First(&u_profile, 1)
	db.Debug().Model(&u_profile).Association("User").Find(&u_profile.User) // "User" is association field
	fmt.Println(u_profile)

	// Query 2: Preload
	var u_profile2 relatetables.UserProfile2
	db.Debug().Preload("User").Find(&u_profile2, 2)
	fmt.Println(u_profile2)

	fmt.Println("----------------------------")
	// // Update
	// var u_profile3 relatetables.UserProfile2
	// db.Preload("User").Find(&u_profile3, 1)
	// fmt.Println(u_profile3)
	// db.Model(&u_profile3.User).Update("addr", "luosaji")
	// fmt.Println(u_profile3)

	// // Del
	var u_profile4 relatetables.UserProfile2
	db.Preload("User").First(&u_profile4, 1)
	fmt.Println(u_profile4)

	db.Delete(&u_profile4.User)
	fmt.Println(u_profile4)

}
