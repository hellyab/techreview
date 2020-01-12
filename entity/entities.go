package entity

// Article entity
type Article struct {
	Name string `gorm:"type:varchar(255);not null"`
	ID   uint
}

type Comment struct {
	CommentBody string `gorm:"type:varchar(255);not null"`
	ID          uint
	UserName    string `gorm:"type:varchar(255);not null"`
}
