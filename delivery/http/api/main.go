package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/hellyab/techreview/delivery/http/handler"
	"github.com/hellyab/techreview/question/repository"
	"github.com/hellyab/techreview/question/service"
)

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tech_review_test?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	//roleRepo
	//roleSrv
	//some role handler

	questionRepo := repository.NewQuestionGormRepo(dbconn)
	questionSrv := service.NewQuestionService(questionRepo)
	questionHandler := handler.NewQuestionHandler(questionSrv)

	router := httprouter.New()

	router.GET("/questions", questionHandler.GetQuestions)
	router.GET("/questions/:id", questionHandler.GetQuestion)
	router.POST("/question", questionHandler.PostQuestion)
	router.PUT("/questions/:id", questionHandler.PutQuestion)
	router.DELETE("/questions/:id", questionHandler.DeleteQuestion)

	handler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	}).Handler(router)
	http.ListenAndServe(":8181", handler)

}
