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

func (as *ArticleService) GetArticle(id uint) (*entity.Article, []error) {
	article, errs := as.articleRepo.GetArticle(id) // here article is pointer btw, it is accesing from the one we did in the repository layer

	if len(errs) > 0 {
		return nil, errs
	}
	return article, nil
}

func (as *ArticleService) PostArticle(article *entity.Article) (*entity.Article, []error) {

	art, errs := as.articleRepo.PostArticle(article)
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs
}

func (as *ArticleService) DeleteArticle(id uint) (*entity.Article, []error) {

	art, errs := as.articleRepo.DeleteArticle(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs

}
