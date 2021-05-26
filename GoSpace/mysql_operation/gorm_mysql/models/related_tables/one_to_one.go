package relatedtables

// Belong To 包含
type User1 struct {
	Id   int
	Name string
	Age  int
	Addr string
}

type UserProfile1 struct {
	Id    int
	Pic   string
	CPic  string
	Phone string
	UId   int   // uid
	User  User1 `gorm:"foreignKey:UId;references:Id"` // 关联关系
	//UserID int  // 默认关联字段为Id
}

// Has one 包含
type User2 struct {
	Id   int
	Name string
	Age  int
	Addr string
	PID  int
}

type UserProfile2 struct {
	Id    int
	Pic   string
	CPic  string
	Phone string
	User  User2 `gorm:"foreignkey:PID"` // 默认关联字段为Id
	// User  User2 `gorm:"foreignkey:PID;references:Id"` // 关联关系
}
