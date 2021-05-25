package main

import (
	dbsource "comegorm/db_source"
	"comegorm/models"
)

var db = dbsource.Db

func main() {
	// db.Migrator().DropTable(&models.User{})
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})
}
