package repository

import (
	"github.com/hellyab/techreview/answer"
	"github.com/hellyab/techreview/entities"
	"github.com/jinzhu/gorm"
)

//AnswerGormRepo implements answer.AnswerRepository interface
type AnswerGormRepo struct {
	conn *gorm.DB
}

//NewAnswerGormRepo returns new object of AnswerGormRepo
func NewAnswerGormRepo(db *gorm.DB) answer.AnswerRepository {
	return &AnswerGormRepo{conn: db}
}

//Answers returns all user answers stored in the database
func (ansRepo *AnswerGormRepo) Answers() ([]entities.Answer, []error) {
	ans := []entities.Answer{}
	errs := ansRepo.conn.Find(&ans).GetErrors()
	if len(errs) > 0 {
		
		return nil, errs
	}
	return ans, errs
}

//Answer returns a user answer stored in the database which has the given id
func (ansRepo *AnswerGormRepo) Answer(id string) (*entities.Answer, []error) {
	qstn := entities.Answer{}
	errs := ansRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	// errs := ansRepo.conn.First(&qstn, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &qstn, errs
}

//UpdateAnswer updates a given answer in the database
func (ansRepo *AnswerGormRepo) UpdateAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn := answer
	errs := ansRepo.conn.Save(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//DeleteAnswer deletes a answer with a given id from the database
func (ansRepo *AnswerGormRepo) DeleteAnswer(id string) (*entities.Answer, []error) {
	qstn, errs := ansRepo.Answer(id)
	if len(errs) > 0 {
		return nil, errs
	}
	// errs := ansRepo.conn.Where("id = ?", id).First(&qstn).GetErrors()
	errs = ansRepo.conn.Delete(qstn).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//StoreAnswer stores a given answer in the database
func (ansRepo *AnswerGormRepo) StoreAnswer(answer *entities.Answer) (*entities.Answer, []error) {
	qstn := answer
	errs := ansRepo.conn.Create(qstn).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return qstn, errs
}

//QuestionAnswers returns answers for a question
func (ansRepo *AnswerGormRepo) QuestionAnswers(question *entities.Question, answer *entities.Answer) ([]entities.Answer, []error) {
	ans := []entities.Answer{}
	errs := ansRepo.conn.Find(&ans).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return ans, errs
}
