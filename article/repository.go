package article

import "github.com/hellyab/techreview/entity"

type ArticleRepository interface {
	Articles() ([]entity.Article, []error)
}
