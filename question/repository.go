package question

import (
	"github.com/hellyab/techreview/entities"
)

//QuestionRepository represents question related operations on the database
type QuestionRepository interface {
	Questions() ([]entities.Question, []error)
	Question(id string) (*entities.Question, []error)
	UpdateQuestion(question *entities.Question) (*entities.Question, []error)
	DeleteQuestion(id string) (*entities.Question, []error)
	StoreQuestion(question *entities.Question) (*entities.Question, []error)
}
