package ddloperate

import (
	dbsource "comegorm/db_source"
	"comegorm/models"
	"fmt"
)

var db = dbsource.Db

func DDLOperation() {
	// Create table for `User`
	db.Migrator().CreateTable(&models.User{})

	// Append "ENGINE=InnoDB" to the creating table SQL for `User`
	// db.Set("gorm:table_options", "ENGINE=InnoDB").Migrator().CreateTable(&models.User{})

	ok := db.Migrator().HasTable(&models.User{})
	// ok := db.Migrator().HasTable("users")
	fmt.Println(ok)

	// Drop table if exists (will ignore or delete foreign key constraints when dropping)
	// db.Migrator().DropTable(&models.User{})
	// db.Migrator().DropTable("users")

	// Rename old table to new table
	// db.Migrator().RenameTable(&User{}, &UserInfo{})
	db.Migrator().RenameTable("users", "user_info")
}
