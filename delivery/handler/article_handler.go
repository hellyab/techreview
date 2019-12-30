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
	articleService article.ArticleService
	tmpl           *template.Template
}

func NewArticleHandler(t *template.Template, artService article.ArticleService) *ArticleHandler {
	return &ArticleHandler{tmpl: t, articleService: artService}
}

func (ah *ArticleHandler) GetArticles(w http.ResponseWriter,
	r *http.Request, _ httprouter.Params) {

	articles, errs := ah.articleService.Articles()

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(articles, "", "\t\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(output)

	// ah.tmpl.ExecuteTemplate(w, "index.html", articles)
	return
}
