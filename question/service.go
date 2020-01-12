package question

import (
	"github.com/hellyab/techreview/entities"
)

//QuestionService represents user - question operations
type QuestionService interface {
	Questions() ([]entities.Question, []error)
	Question(id string) (*entities.Question, []error)
	UpdateQuestion(question *entities.Question) (*entities.Question, []error)
	DeleteQuestion(id string) (*entities.Question, []error)
	StoreQuestion(question *entities.Question) (*entities.Question, []error)
}
