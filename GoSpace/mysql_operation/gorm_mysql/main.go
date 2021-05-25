package main

import (
	dbsource "comegorm/db_source"
	dmloperate "comegorm/dml_operate"
)

var db = dbsource.Db

func main() {
	// db.Migrator().DropTable(&models.User{})
	// db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&models.User{})

	// ddloperate.DDLOperation()
	dmloperate.DMLHandler1()

}
