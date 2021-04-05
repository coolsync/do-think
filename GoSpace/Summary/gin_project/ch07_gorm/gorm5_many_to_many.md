# Many to Many



```go
package relate_tables

// Many To Many
type Article2 struct {
	ID      int
	Title   string
	Content string
	Desc    string
	Tags    []Tag `gorm:"many2many:Article2_tags"` // ;foreignKey:AId;AssociationForeignKey:TId
}

type Tag struct {
	ID   int
	Name string
	Desc string
}

// Many to Many 会在两个 model 中添加一张连接表。

// 例如，您的应用包含了 user 和 language，且一个 user 可以说多种 language，多个 user 也可以说一种 language。

// User 拥有并属于多种 language，`user_languages` 是连接表
// type User3 struct {
// 	gorm.Model
// 	Languages []Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
// 	gorm.Model
// 	Name string
// }

// 反向引用
// type User3 struct {
// 	gorm.Model
// 	Languages []*Language `gorm:"many2many:user_languages;"`
// }

// type Language struct {
// 	gorm.Model
// 	Name string
// 	User3s []*User3 `gorm:"many2many:user_languages;"`
// }

// 重写外键
// type User3 struct {
//     gorm.Model
//     Profiles []Profile `gorm:"many2many:user_profiles;foreignKey:Refer;joinForeignKey:UserReferID;References:UserRefer;JoinReferences:UserRefer"`
//     Refer    uint      `gorm:"index:,unique"`
// }

// type Profile struct {
//     gorm.Model
//     Name      string
//     UserRefer uint `gorm:"index:,unique"`
// }

// 这会创建连接表：user_profiles
//   外键：user_refer_id,，引用：users.refer
//   外键：profile_refer，引用：profiles.user_refer

// CONSTRAINT `fk_user_profiles_profile` FOREIGN KEY (`user_refer`) REFERENCES `profiles` (`user_refer`),
// CONSTRAINT `fk_user_profiles_user3` FOREIGN KEY (`user_refer_id`) REFERENCES `user3` (`refer`)

```





