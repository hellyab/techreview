package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entity"

	"github.com/julienschmidt/httprouter"
)

// handles aricle related http requests
type ArticleHandler struct {
	// ArticleHandler struct has access to ArticleServcie Interface, so as to access repository methods.
	articleService article.ArticleService
	tmpl           *template.Template
}

// Creating an instance of article handler n ArticleHandler implements ArticleService
func NewArticleHandler(t *template.Template, artService article.ArticleService) *ArticleHandler {
	return &ArticleHandler{tmpl: t, articleService: artService}
}

func (ah *ArticleHandler) GetArticles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	articles, errs := ah.articleService.Articles() // get the response (form "repo" -> "service" -> "handler" )

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(articles, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func (ah *ArticleHandler) GetArticle(w http.ResponseWriter,
	r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	article, errs := ah.articleService.GetArticle(uint(id)) // added sample data to fetch by id

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(article, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	return
}

func (ah *ArticleHandler) PostArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	size := r.ContentLength    // set up length accordingly
	body := make([]byte, size) // slice of bytes,
	r.Body.Read(body)          // read request in form of bytes

	article := &entity.Article{} // inti Article struct to put unmarshaled data

	err := json.Unmarshal(body, article) // put the unmarled data of body to aricle struct

	if err != nil {
		// check if error happens , then return status not found 404
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not err

	article, errs := ah.articleService.PostArticle(article) // pass the unmashaled strucl to service and return the article / errs

	// check of errs
	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not errs

	p := fmt.Sprintf("/tech/articles/%d", article.ID) // set up url
	w.Header().Set("Location", p)                     // change url location to /tech/aricles/id
	w.WriteHeader(http.StatusCreated)
	return

}

// Delete article handler

func (ah *ArticleHandler) DeleteArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id")) // get the id from url params

	// check for errs
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not errs

	_, errs := ah.articleService.DeleteArticle(uint(id)) // delete the aricle n if successful ignore the deleted aricle

	// check if errs happened

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not errs

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	return
}

// UpdateArticle updates given article
func (ah *ArticleHandler) UpdateArticle(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	id, err := strconv.Atoi(params.ByName("id")) // get the id form the url prams
	// check for errs
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	// if the id exists in params

	article, errs := ah.articleService.GetArticle(uint(id)) // find an aricle with that id using previous implemneted method getArticle()

	// check for errs

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if the article exists

	size := r.ContentLength    // get the appropiate length for slice
	body := make([]byte, size) // create slice of bytes with length

	r.Body.Read(body) // put read body from request  in the body

	json.Unmarshal(body, &article) // unmarshla the body json, n put it to aricle struct

	article, errs = ah.articleService.UpdateArticle(article) // outsource the aricle to be updated

	// check for errs

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if successfululy udpated

	updatedArticle, err := json.MarshalIndent(article, "", "\t") // get the updated comment n parse it json form

	// check if any errs during marshaling
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if no errs

	w.Header().Set("Content-Type", "application/json")
	w.Write(updatedArticle) //write the updated article
	return
}
