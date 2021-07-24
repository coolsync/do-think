package models

import "gorm.io/gorm"

type GormModel struct {
	gorm.Model
	Name string
}

// Custom table name
// func (GormModel) TableName() string {
// 	return "rename_table_name"
// }