package main

import (
	"recome/db_interface"
	_ "recome/db_source"
)

// var p = fmt.Println
// var db = db_source.Db

func main() {
	db_interface.DBInter1()
}
