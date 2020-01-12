package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/hellyab/techreview/comment"
	"github.com/hellyab/techreview/entity"

	"github.com/julienschmidt/httprouter"
)

// handles aricle related http requests
type CommentHandler struct {
	// CommentHandler struct has access to CommentServcie Interface, so as to access repository methods.
	commentService comment.CommentService
	tmpl           *template.Template
}

// Creating an instance of Comment handler n CommentHandler implements CommentService
func NewCommentHandler(t *template.Template, cmtService comment.CommentService) *CommentHandler {
	return &CommentHandler{tmpl: t, commentService: cmtService}
}

func (ch *CommentHandler) GetComments(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	comments, errs := ch.commentService.Comments() // get the response (form "repo" -> "service" -> "handler" )

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func (ch *CommentHandler) GetComment(w http.ResponseWriter,
	r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	// get the response (form "repo" -> "service" -> "handler" )
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comments, errs := ch.commentService.Comment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comments, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return

}

func (ch *CommentHandler) UpdateComment(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {
	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	comment := &entity.Comment{}
	sampleComment := &entity.Comment{
		ID:          1,
		CommentBody: "changed",
		UserName:    "test user name",
	}
	err := json.Unmarshal(body, sampleComment)

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.commentService.UpdateComment(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	p := fmt.Sprintf("/tech/comments/update/%d", comment.ID)
	w.Header().Set("Location", p)
	w.WriteHeader(http.StatusCreated)
	return
}

func (ch *CommentHandler) DeleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, errs := ch.commentService.DeleteComment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}
func (ch *CommentHandler) PutComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	comment, errs := ch.commentService.Comment(uint(id))

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	l := r.ContentLength

	body := make([]byte, l)

	r.Body.Read(body)

	json.Unmarshal(body, &comment)

	comment, errs = ch.commentService.UpdateComment(comment)

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(comment, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)
	return
}
