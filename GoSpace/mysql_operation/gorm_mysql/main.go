package main

import (
	dbsource "comegorm/db_source"
	relatedtables "comegorm/models/related_tables"
)

var db = dbsource.Db

func main() {
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	// ddloperate.DDLOperation()
	// dml.DMLHandler1()

	// db.AutoMigrate(&models.User{}, &models.GormModel{}, &models.UserInfo{})

	// one to one belong to
	// db.AutoMigrate(&relatedtables.User1{}, &relatedtables.UserProfile1{})

	// has one, 有外键的先迁移
	// db.AutoMigrate(&relatedtables.UserProfile2{},&relatedtables.User2{})

	// one to many
	// db.AutoMigrate(&relatedtables.UserInfo{}, &relatedtables.CreditCard{})

	// many to many
	db.AutoMigrate(&relatedtables.Article{}, &relatedtables.Tag{})

}
