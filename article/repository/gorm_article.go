package repository

import (
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entity"
	"github.com/jinzhu/gorm"
)

type ArticleGormRepo struct {
	conn *gorm.DB
}

// ArticleGormRepo implements article.ArticleRepository
func NewArticleGormRepo(db *gorm.DB) article.ArticleRepository {
	// so, indicating the return type in function, while returing a struct, makes that struct to implement the returned interface
	return &ArticleGormRepo{conn: db}

}

// gets all articles
func (aRepo *ArticleGormRepo) Articles() ([]entity.Article, []error) {
	articles := []entity.Article{}                 //intilize an array of Article entities, so it will be used as a model for GROM
	errs := aRepo.conn.Find(&articles).GetErrors() // growm implementation of finiding array of articles

	if len(errs) > 0 { // if there are erros return nil for data requested and also return array of errs
		return nil, errs
	}
	return articles, errs
}

// gets article by id
func (aRepo *ArticleGormRepo) GetArticle(id uint) (*entity.Article, []error) {
	article := entity.Article{}
	errs := aRepo.conn.First(&article, id).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &article, errs
}

// UpdateArticle: updates article
func (aRepo *ArticleGormRepo) UpdateArticle(article *entity.Article) (*entity.Article, []error) {
	art := article
	errs := aRepo.conn.Save(art).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs

}

// adds new article to db

func (aRepo *ArticleGormRepo) PostArticle(article *entity.Article) (*entity.Article, []error) {
	art := article
	errs := aRepo.conn.Create(art).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs
}

// delete comment from db

func (aRepo *ArticleGormRepo) DeleteArticle(id uint) (*entity.Article, []error) {

	art, errs := aRepo.GetArticle(id) // find the article from the db

	// check for potential errs
	if len(errs) > 0 {
		return nil, errs
	}

	// if not errs finding the article i.e the aricle exists

	errs = aRepo.conn.Delete(art, id).GetErrors() // create conn to db, and delete that aricles with id

	// check for errs

	if len(errs) > 0 {
		return nil, errs
	}

	// if not errs deleting

	return art, errs // return the deleted article

}
