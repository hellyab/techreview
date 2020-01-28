package service

import (
	"github.com/hellyab/techreview/article"
	"github.com/hellyab/techreview/entities"
)

// ArticleService has acces to AricleRepository interface
type ArticleService struct {
	// AricleService struct has access to AricleRepository Interface, so as to access repository methods.
	articleRepo article.ArticleRepository
}

//NewArticleService creates a new AriclesService pointer with access to repository
// AtricleService implements article.service interface
func NewArticleService(ArtRepo article.ArticleRepository) article.ArticleService {
	return &ArticleService {articleRepo: ArtRepo}
}

// Articles outsouces to repository
func (as *ArticleService) Articles() ([]entities.Article, []error) {
	articles, errs := as.articleRepo.Articles()
	if len(errs) > 0 {
		return nil, errs
	}

	return articles, nil
}

//GetArticle outsouces to repository
func (as *ArticleService) GetArticle(id string) (*entities.Article, []error) {
	article, errs := as.articleRepo.GetArticle(id) // here article is pointer btw, it is accesing from the one we did in the repository layer

	if len(errs) > 0 {
		return nil, errs
	}
	return article, nil
}

//PostArticle outsources to repository
func (as *ArticleService) PostArticle(article *entities.Article) (*entities.Article, []error) {

	art, errs := as.articleRepo.PostArticle(article)
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs
}

// DeleteArticle outsouces to repository
func (as *ArticleService) DeleteArticle(id string) (*entities.Article, []error) {

	art, errs := as.articleRepo.DeleteArticle(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return art, errs

}

//UpdateArticle outsources to Repository
func (as *ArticleService) UpdateArticle(article *entities.Article) (*entities.Article, []error) {

	art, errs := as.articleRepo.UpdateArticle(article) // pass the enity from handler to repository

	if len(errs) > 0 {
		return nil, errs
	}

	return art, errs
}

func (as *ArticleService) RateArticle(articleRatings *entities.ArticleRatings) {
	as.articleRepo.RateArticle(articleRatings)
}

func (as *ArticleService) ArticleRateCount(articleId string) int{
	return as.articleRepo.ArticleRateCount(articleId)
}