package entities

import(
	 "time"
)
//User represents application user and has a string ID (which is still being reconsidered)
type User struct {
	ID string
	Username string
	FirstName string
	MiddleName string
	LastName string
	Email string
	Password string
	Interests []Topic
	Privileged bool
}

//Topic represents an area of topic for articles and user's interests. It has a unique string ID and a Name
type Topic struct{
	ID string
	Name string
}

//Article represents a post by a user. It has a unique ID
type Article struct{
	ID string
	AuthorName string
	Content string //TODO needs to be looked into
	Topic []Topic
	AverageRating float32
	NumberOfRatings uint
	PostedAt time.Time
}

//Comment represents comment on article. It has article id and its own unique id.
type Comment struct{
	ID string
	AuthorName string
	Content string //TODO needs to be looked into
	ArticleID string
	PostedAt time.Time
}

//Question represents question asked on the platform. It has its own id and contains the inquirer's id
type Question struct{
	ID string
	InquirerID string
	Inquiry string
}

//Answer represents answer to a question. It has its own id, question's id and the replier's id  
type Answer struct{
	ID string 
	QuestionID string
	ReplierID string
	Answer string
}

