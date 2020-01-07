package repository

import (
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entity"
	"github.com/jinzhu/gorm"
)

type ArticleGormRepo struct {
	conn *gorm.DB
}

func NewArticleGormRepo(db *gorm.DB) article.ArticleRepository {
	return &ArticleGormRepo{conn: db}

}

// ArticleGormRepo implements article.ArticleRepository
func (aRepo *ArticleGormRepo) Articles() ([]entity.Article, []error) {
	articles := []entity.Article{}                 // i didnt get {} , guess it's indicating struct
	errs := aRepo.conn.Find(&articles).GetErrors() // growm implementation

	if len(errs) > 0 { // if there are erros return nil for data requested and also return array of errs
		return nil, errs
	}
	return articles, errs
}
