package models

import (
	"gorm.io/gorm"
)

type GormModel struct {
	gorm.Model
	Name string
}

// type Model struct {
// 	ID        uint      `gorm:"primarykey"`
// 	CreatedAt time.Time // 用于存储记录的创建时间
// 	UpdatedAt time.Time // 用于存储记录的修改时间
// 	DeletedAt DeletedAt `gorm:"index"` // 用于存储记录的删除时间
// }
