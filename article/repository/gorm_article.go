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

func (aRepo *ArticleGormRepo) Articles() ([]entity.Article, []error) {
	articles := []entity.Article{}                 //intilize an array of Article entities, so it will be used as a model for GROM
	errs := aRepo.conn.Find(&articles).GetErrors() // growm implementation of finiding array of articles

	if len(errs) > 0 { // if there are erros return nil for data requested and also return array of errs
		return nil, errs
	}
	return articles, errs
}
