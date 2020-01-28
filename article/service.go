package article

import (
	"github.com/hellyab/techreview/entities"
)

//ArticleService specifies article related actions
type ArticleService interface {
	Articles() ([]entities.Article, []error)
	GetArticle(id string) (*entities.Article, []error)
	// ArticlesByAUser (user *entities.User) ([]*entities.Article, []error)
	// ArticlesOnATopic (topics []*entities.Topic) ([]*entities.Article, []error)
	PostArticle(article *entities.Article) (*entities.Article, []error)
	UpdateArticle(article *entities.Article) (*entities.Article, []error)
	DeleteArticle(id string) (*entities.Article, []error)
	RateArticle(articleRatings *entities.ArticleRatings)
	ArticleRateCount(articleId string) int
	SearchArticle(searchKey string) []entities.Article
}
