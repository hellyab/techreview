package answer

import (
	"github.com/hellyab/techreview/entities"
)

//AnswerService represents user - Answer operations
type AnswerService interface {
	Answers() ([]entities.Answer, []error)
	Answer(id string) (*entities.Answer, []error)
	//QuestionAnswers(question *entities.Question, answer *entities.Answer) ([]entities.Answer, []error)
	UpdateAnswer(Answer *entities.Answer) (*entities.Answer, []error)
	DeleteAnswer(id string) (*entities.Answer, []error)
	StoreAnswer(answer *entities.Answer) (*entities.Answer, []error)
	AnswersByQuestionId(questionId string)([]entities.AnswersByQuesId, []error)
	UpVoteAnswer(answerUpvote *entities.AnswerUpvote)
	UpVoteCount(answerId string) int
}
