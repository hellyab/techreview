package handler

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/hellyab/techreview/article"

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
	r *http.Request, _ httprouter.Params) {

	article, errs := ah.articleService.GetArticle(1) // added sample data to fetch by id

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
