package article

import "github.com/hellyab/techreview/entity"

type ArticleService interface {
	Articles() ([]entity.Article, []error) // gets array of Article structs or reutrns array of errors
}
