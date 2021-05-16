package main

import (
	"recome/db_interface"
	_ "recome/db_source"
)

// var p = fmt.Println
// var db = db_source.Db

func main() {
	// First, FirstOrCreate, Last, Take, Find
	// db_interface.DBInter1()

	// Where, 
	db_interface.DBInter2()
}
