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
