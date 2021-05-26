package relatedtables

import "gorm.io/gorm"

type Article struct {
	ID int `gorm:"primaryKey"`
	// AID     int `gorm:"primaryKey"`
	Title   string
	Content string
	Desc    string
	Tags    []Tag `gorm:"many2many:article_tags"`
	// Tags    []Tag `gorm:"many2many:article_tags;foreignKey:AID;References:TID"`

}

type Tag struct {
	ID   int `gorm:"primaryKey"`
	// TID   int `gorm:"primaryKey"`
	Name string
	Desc string
}

// 没加 `gorm:"primaryKey"`：
// Error 1075: Incorrect table definition; there can be only one auto column and it must be defined as a key

type User3 struct {
	gorm.Model
	Profiles []Profile3 `gorm:"many2many:user3_profile3s;foreignKey:Refer;References:User3Refer"`
	// Profiles []Profile3 `gorm:"many2many:user3_profile3s;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:ProfileRefer"`
	Refer uint
}

type Profile3 struct {
	gorm.Model
	Name      string
	UserRefer uint
}

// Which creates join table: user_profiles
//   foreign key: user_refer_id, reference: users.refer
//   foreign key: profile_refer, reference: profiles.user_refer
