package service

import (
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/question"
)

//QuestionService implements uestion.Service
type QuestionService struct {
	questionRepo question.QuestionRepository
}

// NewQuestionService returns a new QuestionService object
func NewQuestionService(quesRepo question.QuestionRepository) question.QuestionService {
	return &QuestionService{questionRepo: quesRepo}
}

//Questions returns all stored questions
func (qs *QuestionService) Questions() ([]entities.Question, []error) {
	qstns, errs := qs.questionRepo.Questions()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstns, errs
}

//Question retrieves stored question given its id
func (qs *QuestionService) Question(id string) (*entities.Question, []error) {
	qstn, errs := qs.questionRepo.Question(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//UpdateQuestion updates a given question
func (qs *QuestionService) UpdateQuestion(question *entities.Question) (*entities.Question, []error) {
	qstn, errs := qs.questionRepo.UpdateQuestion(question)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//DeleteQuestion deletes a given question
func (qs *QuestionService) DeleteQuestion(id string) (*entities.Question, []error) {
	qstn, errs := qs.questionRepo.DeleteQuestion(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//StoreQuestion stores a given question
func (qs *QuestionService) StoreQuestion(question *entities.Question) (*entities.Question, []error) {
	qstn, errs := qs.questionRepo.StoreQuestion(question)
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}
