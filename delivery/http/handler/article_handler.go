package handler

import (
	"encoding/json"
	"fmt"
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entities"
	"net/http"


	"github.com/julienschmidt/httprouter"
)

//ArticleHandler handles aricle related http requests
type ArticleHandler struct {
	// ArticleHandler struct has access to ArticleServcie Interface, so as to access repository methods.
	articleService article.ArticleService
}

//NewArticleHandler creates an instance of article handler n ArticleHandler implements ArticleService
func NewArticleHandler(artService article.ArticleService) *ArticleHandler {
	return &ArticleHandler{articleService: artService}
}

//GetArticles handler get requests for articles
func (ah *ArticleHandler) GetArticles(w http.ResponseWriter,
	_ *http.Request, _ httprouter.Params) {

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
	_, _ = w.Write(output)

	return
}

//GetArticle handles to get an article
func (ah *ArticleHandler) GetArticle(w http.ResponseWriter,
	_ *http.Request, params httprouter.Params) {

	id :=params.ByName("id")

	art, errs := ah.articleService.GetArticle(id) // added sample handler to fetch by id

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	output, err := json.MarshalIndent(art, "", "\t")

	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(output)

	return
}


//PostArticle handles post methods on articles
func (ah *ArticleHandler) PostArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	size := r.ContentLength    // set up length accordingly
	body := make([]byte, size) // slice of bytes,
	_, _ = r.Body.Read(body)   // read request in form of bytes

	art := &entities.Article{} // inti Article struct to put unmarshaled handler

	err := json.Unmarshal(body, art) // put the unmarled handler of body to aricle struct

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("error while unmarshing article")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not err

	art, errs := ah.articleService.PostArticle(art) // pass the unmashaled strucl to service and return the article / errs

	// check of errs
	if len(errs) > 0 {
		fmt.Println("erro while storing article")
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if not errs
// set up url                    // change url location to /tech/aricles/id
//	art, errs = ah.articleService.GetArticle(art.ID)
//
//	if len(errs) > 0 {
//		w.Header().Set("Content-Type", "application/json")
//		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
//		return
//	}

	output, err := json.MarshalIndent(art, "", "\t")

	if err!=nil{
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusCreated)
	_, _ = w.Write(output)
	return

}

// DeleteArticle handles delete request on articles
func (ah *ArticleHandler) DeleteArticle(w http.ResponseWriter, _ *http.Request, params httprouter.Params) {

	id:= params.ByName("id") // get the id from url params


	_, errs := ah.articleService.DeleteArticle(id) // delete the aricle n if successful ignore the deleted aricle

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

	id := params.ByName("id") // get the id form the url prams

	art, errs := ah.articleService.GetArticle(id) // find an aricle with that id using previous implemneted method getArticle()

	// check for errs

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if the article exists

	l:= r.ContentLength    // get the appropiate length for slice
	body := make([]byte, l) // create slice of bytes with length

	_, _ = r.Body.Read(body) // put read body from request  in the body

	_ = json.Unmarshal(body, &art) // unmarshlal the body json, n put it to aricle struct

	art, errs = ah.articleService.UpdateArticle(art) // outsource the aricle to be updated

	// check for errs

	if len(errs) > 0 {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if successfululy udpated

	updatedArticle, err := json.MarshalIndent(art, "", "\t") // get the updated comment n parse it json form

	// check if any errs during marshaling
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	// if no errs

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(updatedArticle) //write the updated article
	return
}

func (ah *ArticleHandler) RateArticle(w http.ResponseWriter, r *http.Request, _ httprouter.Params){

	l := r.ContentLength
	body := make([]byte, l)
	r.Body.Read(body)
	artRating := &entities.ArticleRatings{}

	err := json.Unmarshal(body, artRating)
	fmt.Println("successfully read the body and assinged to the artRating struct")
	if err != nil {
		fmt.Println("error while unmarshing the artRating", err)
		w.Header().Set("Content-Type", "application/json")
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		//return
	}

	ah.articleService.RateArticle(artRating)
	fmt.Println("success form article rating handler")
	return
}


func (ah *ArticleHandler) ArticleRateCount(w http.ResponseWriter, _ *http.Request, params httprouter.Params){
	ansId := params.ByName("artId")

	rateCount := ah.articleService.ArticleRateCount(ansId)

	output, err := json.MarshalIndent(rateCount, "", "\t")

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

