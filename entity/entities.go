package entity

import (
	"time"
)

//User represents application user and has a string ID (which is still being reconsidered)
type User struct {
	ID         string `gorm:"type:varchar(255);not null"`
	Username   string `gorm:"type:varchar(255);not null"`
	FirstName  string `gorm:"type:varchar(255);not null"`
	MiddleName string `gorm:"type:varchar(255);not null"`
	LastName   string `gorm:"type:varchar(255);not null"`
	Email      string `gorm:"type:varchar(255);not null"`
	Password   string `gorm:"type:varchar(255);not null"`
	Interests  []Topic
	Privileged bool
}

//Topic represents an area of topic for articles and user's interests. It has a unique string ID and a Name
type Topic struct {
	ID   string
	Name string `gorm:"type:varchar(255);not null"`
}

//Article represents a post by a user. It has a unique ID
type Article struct {
	ID              uint
	AuthorName      string `gorm:"type:varchar(255);not null"`
	Content         string `gorm:"type:varchar(255);not null"` //TODO needs to be looked into
	Topic           []Topic
	AverageRating   float32
	NumberOfRatings uint
	PostedAt        time.Time
}

//Comment represents comment on article. It has article id and its own unique id.
type Comment struct {
	ID         string `gorm:"type:varchar(255);not null"`
	AuthorName string `gorm:"type:varchar(255);not null"`
	Content    string `gorm:"type:varchar(255);not null"` //TODO needs to be looked into
	ArticleID  string
	PostedAt   time.Time
}

//Question represents question asked on the platform. It has its own id and contains the inquirer's id
type Question struct {
	ID         string `json:"ID"`
	InquirerID string `gorm:"column:inquirer" json:"InquirerID"`
	Inquiry    string `json:"Inquiry"`
}

//TableName changes the name of the table for gorm
func (Question) TableName() string {
	return "question"
}

//Answer represents answer to a question. It has its own id, question's id and the replier's id
type Answer struct {
	ID         string `gorm:"type:varchar(255);not null"`
	QuestionID string `gorm:"type:varchar(255);not null"`
	ReplierID  string `gorm:"type:varchar(255);not null"`
	Answer     string `gorm:"type:varchar(255);not null"`
}
