package handler

import (
	"encoding/json"
	"net/http"

	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/question"
	"github.com/julienschmidt/httprouter"
)

//QuestionHandler handles question related requests from a user
type QuestionHandler struct {
	questionService question.QuestionService
}

//NewQuestionHandler (for user) returns new QuestionHandler object
func NewQuestionHandler(qstnService question.QuestionService) *QuestionHandler {
	return &QuestionHandler{questionService: qstnService}
}

//GetQuestions (for us er)handles GET /questions request
func (qh *QuestionHandler) GetQuestions(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	questions, errs := qh.questionService.Questions()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(questions, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "applcation/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}

//GetQuestion handles GET questions/:id request
func (qh *QuestionHandler) GetQuestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	question, errs := qh.questionService.Question(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(question, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//PostQuestion handles POST question request
func (qh *QuestionHandler) PostQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	question := &entities.Question{}

	err := json.Unmarshal(body, question)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	question, errs := qh.questionService.StoreQuestion(question)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := "questions/" + question.ID

	w.Header().Set("Location", p)
	http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
	return
}

//PutQuestion handles PUT questions/:id request
func (qh *QuestionHandler) PutQuestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	question, errs := qh.questionService.Question(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &question)

	question, errs = qh.questionService.UpdateQuestion(question)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(question, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//DeleteQuestion handles DELETE questions/:id request
func (qh *QuestionHandler) DeleteQuestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deletedQuestion, errs := qh.questionService.DeleteQuestion(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// log.Println(deletedQuestion)
	output, err := json.MarshalIndent(deletedQuestion, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// log.Println(output)
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusNoContent)
	w.Write(output)
	return
}
