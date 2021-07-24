package relate_tables

// belong to
type User struct {
	ID   int
	Name string
	Age  int
	Addr string
	PID int
}

type UserProfile struct {
	ID    int
	Pic   string
	CPic  string
	Phone string
	User  User `gorm:"foreignKey:PID;references:ID"`
}

// delete table, 
// 1 del users
// 2 del user_profiles 
