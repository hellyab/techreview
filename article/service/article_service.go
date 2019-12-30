package service

import (
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entity"
)

type ArticleService struct {
	articleRepo article.ArticleRepository
}

func NewArticleService(ArtRepo article.ArticleRepository) article.ArticleService {
	return &ArticleService{articleRepo: ArtRepo}
}

// AtricleService implements article.service interface

func (as *ArticleService) Articles() ([]entity.Article, []error) {
	articles, errs := as.articleRepo.Articles()
	if len(errs) > 0 {
		return nil, errs
	}

	return articles, nil
}
