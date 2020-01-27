package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hellyab/techreview/entities"
	"github.com/hellyab/techreview/question"
	"net/http"


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
		fmt.Println("errors while puting the json in to the struct")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("successfully unmarshed the json body")

	question, errs := qh.questionService.StoreQuestion(question)

	if errs != nil{
		fmt.Println("error while stroing the question")
	}

	output, err := json.MarshalIndent(question, "", "\t")

	if err != nil {
		fmt.Println("error while marhsing the srct to json")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("we got the output as json")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
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
	fmt.Println("we got the output as json")

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


func (qh *QuestionHandler) FollowQuestion(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	questionFollow := &entities.QuestionFollow{}

	err := json.Unmarshal(body, questionFollow)

	if err != nil {
		fmt.Println("errors while puting the json in to the questionFollow struct")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("successfully unmarshed the json body")

	questionExists:= qh.questionService.FollowQuestion(questionFollow.QuestionID, questionFollow.UserID)

	if !questionExists{
		fmt.Println("quesition dosen't exist", questionExists)
	}

	output, err := json.MarshalIndent(questionExists, "", "\t")

	if err != nil {
		fmt.Println("error while marhsing the questionExists boolean to json")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("we got the output as json, the question exist boolean")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}


func (qh *QuestionHandler) FollowedByUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	questionFollow := &entities.QuestionFollow{}

	err := json.Unmarshal(body, questionFollow)

	if err != nil {
		fmt.Println("errors while puting the json in to the questionFollow struct")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("successfully unmarshed the json body")

	questionExists:= qh.questionService.FollowedByUser(questionFollow)


	output, err := json.MarshalIndent(questionExists, "", "\t")

	if err != nil {
		fmt.Println("error while marhsing the questionExists boolean to json")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	fmt.Println("we got the output as json, the question exist boolean")

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}


func (qh *QuestionHandler) FollowCount(w http.ResponseWriter, r *http.Request, params httprouter.Params){

	quesId := params.ByName("quesId")

	quesCount := qh.questionService.FollowCount(quesId)

	if quesCount == -1 {
		fmt.Println("not working")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	output, err := json.MarshalIndent(quesCount, "", "\t")

	if err != nil {
		fmt.Println("error while marshaling")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return

}