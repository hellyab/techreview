package article

import "github.com/hellyab/techreview/entity"

type ArticleService interface {
	Articles() ([]entity.Article, []error) // gets array of Article structs or reutrns array of errors
	GetArticle(id uint) (*entity.Article, []error)
	// ArticlesByAUser (user *entity.User) ([]*entity.Article, []error)
	// ArticlesOnATopic (topics []*entity.Topic) ([]*entity.Article, []error)
	PostArticle(article *entity.Article) (*entity.Article, []error)
	// UpdateArticle(article *entity.Article) (*entity.Article, []error)
	DeleteArticle(id uint) (*entity.Article, []error)
}
