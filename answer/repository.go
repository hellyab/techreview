package answer

import (
	"github.com/hellyab/techreview/entities"
)

//AnswerRepository represents user - Answer operations
type AnswerRepository interface {
	Answers() ([]entities.Answer, []error)
	Answer(id string) (*entities.Answer, []error)
	UpdateAnswer(Answer *entities.Answer) (*entities.Answer, []error)
	DeleteAnswer(id string) (*entities.Answer, []error)
	StoreAnswer(answer *entities.Answer) (*entities.Answer, []error)
	AnswersByQuestionId(questionId string)([]entities.AnswersByQuesId, []error)
}
