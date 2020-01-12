package main

import (
	"net/http"
	"text/template"

	"github.com/hellyab/techreview"
	"github.com/hellyab/techreview/article/repository"
	"github.com/hellyab/techreview/article/service"

	"github.com/hellyab/techreview/delivery/handler"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/julienschmidt/httprouter"
)

func main() {
	// connect to db
	dbconn, err := gorm.Open(techreview.DBConnfigurations.DatabaseName,
		techreview.DBConnfigurations.ConnString)

	// check for db err
	if err != nil {
		panic(err)
	}
	defer dbconn.Close()

	// initilize the ORM
	articleRepo := repository.NewArticleGormRepo(dbconn)

	articleSrv := service.NewArticleService(articleRepo)
	tmpl := template.Must(template.ParseGlob("../ui/templates/*"))

	articleHandler := handler.NewArticleHandler(tmpl, articleSrv)

	router := httprouter.New()

	router.GET("/tech/articles", articleHandler.GetArticles)
	router.GET("/tech/articles/:id", articleHandler.GetArticle)
	router.POST("/tech/articles", articleHandler.PostArticle)
	router.DELETE("/tech/articles/:id", articleHandler.DeleteArticle)
	http.ListenAndServe("localhost:8181", router)
}
