package article

import "github.com/hellyab/techreview/entities"

//ArticleRepository porvides the interface to store entities
type ArticleRepository interface {
	Articles() ([]entities.Article, []error)
	GetArticle(id string) (*entities.Article, []error)
	// ArticlesByAUser (user *entities.User) ([]*entities.Article, []error)
	// ArticlesOnATopic (topics []*entities.Topic) ([]*entities.Article, []error)
	PostArticle(article *entities.Article) (*entities.Article, []error)
	UpdateArticle(article *entities.Article) (*entities.Article, []error)
	DeleteArticle(id string) (*entities.Article, []error)
	RateArticle(articleRatings *entities.ArticleRatings)
	ArticleRateCount(articleId string) int


}
