package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"

	"github.com/hellyab/techreview/delivery/http/handler"

	ansRepo "github.com/hellyab/techreview/answer/repository"
	ansServ "github.com/hellyab/techreview/answer/service"

	questRepo "github.com/hellyab/techreview/question/repository"
	questServ "github.com/hellyab/techreview/question/service"

	commRepo "github.com/hellyab/techreview/comment/repository"
	commServ "github.com/hellyab/techreview/comment/service"

	usRepo "github.com/hellyab/techreview/user/repository"
	usServ "github.com/hellyab/techreview/user/service"

	artRepo "github.com/hellyab/techreview/article/repository"
	artServ "github.com/hellyab/techreview/article/service"
)

//roleRepo
//roleSrv
//some role handler

func main() {
	dbconn, err := gorm.Open("postgres", "postgres://postgres:password@localhost/tech_review_test?sslmode=disable")

	if err != nil {
		panic(err)
	}

	defer dbconn.Close()

	questionRepo := questRepo.NewQuestionGormRepo(dbconn)
	questionSrv := questServ.NewQuestionService(questionRepo)
	questionHandler := handler.NewQuestionHandler(questionSrv)

	answerRepo := ansRepo.NewAnswerGormRepo(dbconn)
	answerSrv := ansServ.NewAnswerService(answerRepo)
	answerHandler := handler.NewAnswerHandler(answerSrv)

	commentRepo := commRepo.NewCommentGormRepo(dbconn)
	commentSrv := commServ.NewCommentService(commentRepo)
	commentHandler := handler.NewCommentHandler(commentSrv)

	userRepo := usRepo.NewUserGormRepo(dbconn)
	userSrv := usServ.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userSrv)

	articleRepo := artRepo.NewArticleGormRepo(dbconn)
	articleSrv := artServ.NewArticleService(articleRepo)
	articleHandler := handler.NewArticleHandler(articleSrv)

	router := httprouter.New()

	router.GET("/questions", questionHandler.GetQuestions)
	router.GET("/questions/:id", questionHandler.GetQuestion)
	router.POST("/questions", questionHandler.PostQuestion)
	router.PUT("/questions/:id", questionHandler.PutQuestion)
	router.DELETE("/questions/:id", questionHandler.DeleteQuestion)

	router.GET("/answers", answerHandler.GetAnswers)
	router.GET("/answers/:id", answerHandler.GetAnswer)
	router.POST("/answers", answerHandler.PostAnswer)
	router.PUT("/answers/:id", answerHandler.PutAnswer)
	router.DELETE("/answers/:id", answerHandler.DeleteAnswer)

	router.GET("/comments", commentHandler.GetComments)
	router.GET("/comments/:id", commentHandler.GetComment)
	router.POST("/comments", commentHandler.UpdateComment)
	router.DELETE("/comments/:id", commentHandler.DeleteComment)
	router.PUT("/comments/:id", commentHandler.PutComment)

	router.GET("/users", userHandler.GetUsers)
	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.AddUser)
	router.DELETE("/users/:id", userHandler.DeleteUser)
	router.PUT("/users/:id", userHandler.UpdateUser)

	router.GET("/articles", articleHandler.GetArticles)
	router.GET("/articles/:id", articleHandler.GetArticle)
	router.POST("/articles", articleHandler.PostArticle)
	router.DELETE("/articles/:id", articleHandler.DeleteArticle)
	router.PUT("/articles/:id", articleHandler.UpdateArticle)

	apiHandler := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "OPTIONS"},
	}).Handler(router)

	http.ListenAndServe(":8181", apiHandler)

}
