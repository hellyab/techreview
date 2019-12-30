package main

import (
	"net/http"
	"text/template"

	"github.com/hellyab/techreview/article/repository"
	"github.com/hellyab/techreview/article/service"
	"github.com/hellyab/techreview/delivery/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/julienschmidt/httprouter"
)

func main() {

	dbconn, err := gorm.Open("postgres",
		"postgres://postgres:Binaman1!@localhost/testdb?sslmode=disable")

	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	articleRepo := repository.NewArticleGormRepo(dbconn)
	articleSrv := service.NewArticleService(articleRepo)
	tmpl := template.Must(template.ParseGlob("../ui/templates/*"))

	articleHandler := handler.NewArticleHandler(tmpl, articleSrv)

	router := httprouter.New()

	router.GET("/v1/articles", articleHandler.GetArticles)

	http.ListenAndServe(":8181", router)
}
