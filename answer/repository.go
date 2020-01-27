<<<<<<< HEAD
package answer

import (
	"github.com/hellyab/techreview/entities"
)

//AnswerRepository represents user - Answer operations
type AnswerRepository interface {
	Answers() ([]entities.Answer, []error)
	Answer(id string) (*entities.Answer, []error)
	QuestionAnswers(question *entities.Question, answer *entities.Answer) ([]entities.Answer, []error)
	UpdateAnswer(Answer *entities.Answer) (*entities.Answer, []error)
	DeleteAnswer(id string) (*entities.Answer, []error)
	StoreAnswer(answer *entities.Answer) (*entities.Answer, []error)
}
=======
package answer

import (
	"github.com/hellyab/techreview/entities"
)

//AnswerRepository represents user - Answer operations
type AnswerRepository interface {
	Answers() ([]entities.Answer, []error)
	Answer(id string) (*entities.Answer, []error)
	QuestionAnswers(question *entities.Question, answer *entities.Answer) ([]entities.Answer, []error)
	UpdateAnswer(Answer *entities.Answer) (*entities.Answer, []error)
	DeleteAnswer(id string) (*entities.Answer, []error)
	StoreAnswer(answer *entities.Answer) (*entities.Answer, []error)
}
>>>>>>> fbe394209bc600b01af6f1f873d27d6f2c253b44
