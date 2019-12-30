package entity

// Article entity
type Article struct {
	Name string `gorm:"type:varchar(255);not null"`
	ID   uint
}
