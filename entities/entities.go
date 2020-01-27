package entities

import (
	"encoding/json"
	"time"
)

//User represents application user and has a string ID (which is still being reconsidered)
type User struct {
	//ID         string `gorm:"type:varchar; default; replace(uuid_generate_v4() :: text, '-', '')`
	Username   string
	FirstName  string
	LastName   string
	Email      string
	Password   string
	RoleID     string
	Interests  json.RawMessage
}

//TableName changes the name of the table
//func (User) TableName() string {
//	return "person"
//}

//Topic represents an area of topic for articles and user's interests. It has a unique string ID and a Name
type Topic struct {
	ID   uint
	Name string
}

//Article represents a post by a user. It has a unique ID
type Article struct {
	ID              string
	AuthorID          string
	Content         json.RawMessage
	Topics          json.RawMessage
	AverageRating   float32
	NumberOfRatings uint
	PostedAt        time.Time
	Review		bool
	NumberOfComments uint

}

//Comment represents comment on article. It has article id and its own unique id.
type Comment struct {
	ID         string
	Writer string
	Content    string //TODO needs to be looked into
	ArticleID  string
	PostedAt   time.Time
	Likes uint
}

//Question represents question asked on the platform. It has its own id and contains the inquirer's id
type Question struct {
	ID         string
	InquirerId string
	Inquiry    string
	AskedAt time.Time
	Follows uint
	NumberOfAnswers uint
	Topics json.RawMessage
}

//Answer represents answer to a question. It has its own id, question's id and the replier's id
type Answer struct {
	ID         string
	QuestionID string
	ReplierId string
	Answer     string
	Votes uint
}

type ArticleRating struct {
	ArticleID string
	UserID string

}

type CommentLike struct{
	ArticleID string
	UserID string
}

type AnswerUpvote struct {
	AnswerID string
	UserID string
}

type QuestionFollow struct{
	QuestionID string
	UserID string
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
	ID    string
	Name  string `gorm:"type:varchar(255)"`
}

type AnswersByQuesId struct{
	Votes int
	Answer string
	AskedByUserName string
	AskedByFirstName string
	AskedByLastName string
	AnswerId string

}