package handler

import (
	"encoding/json"
	"github.com/hellyab/techreview/answer"
	"github.com/hellyab/techreview/entities"
	"net/http"
	


	"github.com/julienschmidt/httprouter"
)

//AnswerHandler handles menu related requests from a user
type AnswerHandler struct {
	answerService answer.AnswerService
}

//NewAnswerHandler (for user) returns new AnswerHandler object
func NewAnswerHandler(qstnService answer.AnswerService) *AnswerHandler {
	return &AnswerHandler{answerService: qstnService}
}

//GetAnswers (for user)handles GET /answers request
func (ah *AnswerHandler) GetAnswers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	answers, errs := ah.answerService.Answers()

	if len(errs) > 0 {

		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(answers, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "applcation/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

}

//GetAnswer handles GET answers/:id request
func (ah *AnswerHandler) GetAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id := ps.ByName("id")

	answer, errs := ah.answerService.Answer(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(answer, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//PostAnswer handles POST answer request
func (ah *AnswerHandler) PostAnswer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	answer := &entities.Answer{}

	err := json.Unmarshal(body, answer)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	answer, errs := ah.answerService.StoreAnswer(answer)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := "answers/" + answer.ID

	w.Header().Set("Location", p)
	http.Error(w, http.StatusText(http.StatusCreated), http.StatusCreated)
	return
}

//PutAnswer handles PUT answers/:id request
func (ah *AnswerHandler) PutAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	answer, errs := ah.answerService.Answer(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &answer)

	answer, errs = ah.answerService.UpdateAnswer(answer)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(answer, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}

//DeleteAnswer handles DELETE answers/:id request
func (ah *AnswerHandler) DeleteAnswer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	deletedAnswer, errs := ah.answerService.DeleteAnswer(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// log.Println(deletedAnswer)
	output, err := json.MarshalIndent(deletedAnswer, "", "\t")

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


func (ah *AnswerHandler) GetAnswersByQuestionId(w http.ResponseWriter, _ *http.Request, params httprouter.Params){


	id := params.ByName("questionId")

	answersByQuesId, errs := ah.answerService.AnswersByQuestionId(id)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(answersByQuesId, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
