# Associations

# One to One



```go
// Belongs To
// belongs to 会与另一个模型建立了一对一的连接。 这种模型的每一个实例都“属于”另一个模型的一个实例。
// 例如，您的应用包含 user 和 company，并且每个 user 都可以分配给一个 company

// // 1
// type User struct {
// 	gorm.Model
// 	Name         string
// 	CompanyRefer int
// 	Company      Company `gorm:"foreignKey:CompanyRefer"`
// }

// type Company struct {
// 	ID   int
// 	Name string
// }

// 2
type User struct {
	ID   int
	Name string
	Age  string
	Addr string
}

type UserProfile struct {
	ID    int
	Pic   string
	CPic  string
	Phone string
	// User  User // relate // 关联关系
	User User `gorm:"foreignKey:UID;references:ID"` // relate // 关联关系
	UID  int
}
```



```go
// Has One
// has one 与另一个模型建立一对一的关联，但它和一对一关系有些许不同。 这种关联表明一个模型的每个实例都包含或拥有另一个模型的一个实例。
type User struct {
	ID   int
	Name string
	Age  string
	Addr string
	UserProfile UserProfile `gorm:"foreignKey:UID;references:ID"`
}

type UserProfile struct {
	ID    int
	Pic   string
	CPic  string
	Phone string
	// User  User // relate // 关联关系
	// User User `gorm:"foreignKey:PID;references:ID"` // relate // 关联关系
	UID int	// foreign key
}

// Error 1826: Duplicate foreign key constraint name 'fk_user_profiles_user'
// Error 1824: Failed to open the referenced table 'user_profiles'
```



```go
// 属于：关系和外键在同一方，有关系的那一方属于另外一个模型
// 包含：关系和外键不在同一方，有关系的那一方包含另外一个有外键的模型
```







# One to One Operate
















