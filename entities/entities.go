package entities

import (
	"encoding/json"
)

//User represents application user and has a string ID (which is still being reconsidered)
type User struct {

	ID         string `json:"ID,omitempty"`
	Username   string
	FirstName  string
	MiddleName string
	LastName   string
	Email      string
	Password   string
	RoleID     uint
	Interests  json.RawMessage

}

//TableName changes the name of the table
//func (User) TableName() string {
//	return "person"
//}

//Topic represents an area of topic for articles and user's interests. It has a unique string ID and a Name
type Topic struct {
	ID   string
	Name string
}

//Article represents a post by a user. It has a unique ID
type Article struct {
	ID              string
	Author          string
	Content         json.RawMessage
	Topics          json.RawMessage
	AverageRating   float32
	NumberOfRatings uint
	PostedAt        time.Time
}

func (Article) TableName() string {
	return "article"

}

//Comment represents comment on article. It has article id and its own unique id.
type Comment struct {
	ID         string
	AuthorName string
	Content    string //TODO needs to be looked into
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
	ID         string
	QuestionID string
	ReplierID  string
	Answer     string
}

//TableName changes the name of the table for gorm
func (Answer) TableName() string {
	return "answer"
}

//Session represents login user session
type Session struct {
	ID         uint
	UUID       string `gorm:"type:varchar(255);not null"`
	Expires    int64  `gorm:"type:varchar(255);not null"`
	SigningKey []byte `gorm:"type:varchar(255);not null"`
}

// Role repesents application user roles
type Role struct {
	ID    uint
	Name  string `gorm:"type:varchar(255)"`
	Users []User
}
