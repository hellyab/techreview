package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"


	ansRepo "github.com/hellyab/techreview/answer/repository"
	ansServ "github.com/hellyab/techreview/answer/service"
	"github.com/hellyab/techreview/delivery/http/handler"
	questRepo "github.com/hellyab/techreview/question/repository"
	questServ "github.com/hellyab/techreview/question/service"
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

	questionRepo := questRepo.NewQuestionGormRepo(dbconn)
	questionSrv := questServ.NewQuestionService(questionRepo)
	questionHandler := handler.NewQuestionHandler(questionSrv)

	answerRepo := ansRepo.NewAnswerGormRepo(dbconn)
	answerSrv := ansServ.NewAnswerService(answerRepo)
	answerHandler := handler.NewAnswerHandler(answerSrv)

	router := httprouter.New()

	router.GET("/questions", questionHandler.GetQuestions)
	router.GET("/questions/:id", questionHandler.GetQuestion)
	router.POST("/question", questionHandler.PostQuestion)
	router.PUT("/questions/:id", questionHandler.PutQuestion)
	router.DELETE("/questions/:id", questionHandler.DeleteQuestion)

	router.GET("/answers", answerHandler.GetAnswers)
	router.GET("/answers/:id", answerHandler.GetAnswer)
	router.POST("/answer", answerHandler.PostAnswer)
	router.PUT("/answers/:id", answerHandler.PutAnswer)
	router.DELETE("/answers/:id", answerHandler.DeleteAnswer)

	apiHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8181", apiHandler)

}
