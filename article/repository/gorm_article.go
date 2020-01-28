package repository

import (
	"fmt"
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entities"
	"github.com/jinzhu/gorm"
)

//ArticleGormRepo has acces to gorm
type ArticleGormRepo struct {
	conn *gorm.DB
}

//NewArticleGormRepo ArticleGormRepo implements article.ArticleRepository
func NewArticleGormRepo(db *gorm.DB) article.ArticleRepository {
	// so, indicating the return type in function, while returing a struct, makes that struct to implement the returned interface
	return &ArticleGormRepo{conn: db}

}

// Articles gets all articles
func (aRepo *ArticleGormRepo) Articles() ([]entities.Article, []error) {
	articles := []entities.Article{}               //intilize an array of Article entities, so it will be used as a model for GROM
	errs := aRepo.conn.Find(&articles).GetErrors() // growm implementation of finiding array of articles

	if len(errs) > 0 { // if there are erros return nil for handler requested and also return array of errs
		return nil, errs
	}
	return articles, errs
}

//GetArticle gets article by id
func (aRepo *ArticleGormRepo) GetArticle(id string) (*entities.Article, []error) {
	article := entities.Article{}
	errs := aRepo.conn.Where("id = ?", id).First(&article).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &article, errs
}

//PostArticle adds new article to db
func (aRepo *ArticleGormRepo) PostArticle(article *entities.Article) (*entities.Article, []error) {
	art := article
	errs := aRepo.conn.Create(art).GetErrors()
	if len(errs) > 0 {
		fmt.Println("error creating article")
		return nil, errs
	}
	return art, errs
}

//DeleteArticle deletes article	 get the aricle sturct pointer n assing to art from db
func (aRepo *ArticleGormRepo) DeleteArticle(id string) (*entities.Article, []error) {

	article, errs := aRepo.GetArticle(id)
	if len(errs) > 0 {
		return nil, errs
	}
	
	// errs := aRepo.conn.Where("id = ?", id).First(&article).GetErrors()
	// // find the article from the db

	// // check for potential errs
	// if len(errs) > 0 {
	// 	return nil, errs
	// }

	// if not errs finding the article i.e the aricle exists

	errs = aRepo.conn.Delete(article).GetErrors() // create conn to db, and delete that aricles with id

	// check for errs

	if len(errs) > 0 {
		return nil, errs
	}

	// if not errs deleting

	return article, errs // return the deleted article

}

//UpdateArticle updates article
func (aRepo *ArticleGormRepo) UpdateArticle(article *entities.Article) (*entities.Article, []error) {

	art := article                           // get the aricle sturct pointer n assing to art
	errs := aRepo.conn.Save(art).GetErrors() // do orm method SAVE
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs
}


func (aRepo *ArticleGormRepo) RateArticle(articleRatings *entities.ArticleRatings) {

	errs := aRepo.conn.Where("article_id = ? AND user_id= ?", articleRatings.ArticleID, articleRatings.UserID).First(&articleRatings).GetErrors()

	if len(errs) > 0 {

		errs := aRepo.conn.Create(articleRatings).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error storing article ratings")
			return
		}
		fmt.Println("stored  article ratings")
		fmt.Println("query worked")
		return

	} else{
		errs := aRepo.conn.Delete(articleRatings).GetErrors()
		if len(errs) > 0 {
			fmt.Println("error while deleting article ratings", errs)
			return
		}
		fmt.Println("deleted answer article ratings successfuly")
		return
	}
}

func (aRepo *ArticleGormRepo) ArticleRateCount(articleId string) int{
	var count int
	var aritlceRating entities.ArticleRatings
	//questionFollows := entities.QuestionFollow{}
	aRepo.conn.Model(&aritlceRating).Where("article_id = ?", articleId).Count(&count)
	// update the count value in articles table
	aRepo.conn.Model(&entities.Article{}).Where("id = ?", articleId).UpdateColumn("number_of_ratings", count)
	// update the average ratings in articles table
	aRepo.conn.Model(&entities.Article{}).Where("id = ?", articleId).UpdateColumn("average_rating", count/5) // it is doing it just the data

	fmt.Println("counted succesfully, with value: ", count)

	return count
}

func (aRepo *ArticleGormRepo) SearchArticle(searchKey string) []entities.Article{

	results := []entities.Article{}

	err := aRepo.conn.Table("articles").Where(" @@ to_tsquery(?)", searchKey).Find(&results).GetErrors()

	if err != nil{
		fmt.Println("error searching articles")

	}

	return results

}