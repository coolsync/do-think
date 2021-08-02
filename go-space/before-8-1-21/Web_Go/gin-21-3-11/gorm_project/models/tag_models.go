package models

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	// Title   string `gorm:"not null;index"`	// idx_articles_title
	// Title   string `gorm:"not null;unique_index"`
	Title   string `gorm:"not null;unique"`
	Content string `gorm:"column:a_content;size:64"`
	Desc    string `gorm:"type:int(11)"`
	Test    string `gorm:"-"`
}
