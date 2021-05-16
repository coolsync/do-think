package relate_tables

// one to many
type User2 struct {
	ID int
	Name string
	Age int
	Addr string
	Articles []Article `gorm:"foreignKey:UID;references:ID"`
}

type Article struct {
	ID int
	Title string
	Content string
	Desc string
	UID int // foreign key
}
