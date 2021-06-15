package relatetables

import "gorm.io/gorm"

// User 有多张 CreditCard，UserID 是外键
type UserInfo struct {
	gorm.Model
	// MemberNumber string
	CreditCards  []CreditCard `gorm:"foreignKey:UserID;references:ID"`
}

type CreditCard struct {
	ID uint
	Number string
	UserID uint
}
