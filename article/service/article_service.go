package service

import (
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entity"
)

type ArticleService struct {
	// AricleService struct has access to AricleRepository Interface, so as to access repository methods.
	articleRepo article.ArticleRepository
}

// AtricleService implements article.service interface
func NewArticleService(ArtRepo article.ArticleRepository) article.ArticleService {
	return &ArticleService{articleRepo: ArtRepo}
}

func (as *ArticleService) Articles() ([]entity.Article, []error) {
	articles, errs := as.articleRepo.Articles()
	if len(errs) > 0 {
		return nil, errs
	}

	return articles, nil
}
