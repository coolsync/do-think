package dmloperate

import (
	dbsource "comegorm/db_source"
	"comegorm/models"
	"fmt"
)

var db = dbsource.Db

func DMLHandler1() {

	// Create
	// db.Create(&models.User{Name: "bob", Age: 30, Addr: "xxx", Pic: "/static/upload/pic.jpg"})

	// Query
	var user1 models.User
	db.First(&user1, 2) // id = 2
	// db.First(&user1, "name=?", "bob")
	fmt.Println("user1: ", user1)

	// Update
	var user2 models.User
	db.First(&user2, 2)
	user2.Name = "paul"
	user2.Age = 20
	db.Save(&user2)

	db.Model(&user2).Update("addr", "paul-yyyy")
	db.Model(&user2).Update("phone", "12345678")

	// Del
	var user3 models.User
	db.First(&user3, 4)
	db.Delete(&user3)

	db.Where("name", "mark").Delete(&models.User{})
}
