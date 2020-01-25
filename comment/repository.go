package comment

import "github.com/hellyab/techreview/entities"

type CommentRepository interface {
	Comments() ([]entities.Comment, []error)
	Comment(id uint) (*entities.Comment, []error)
	UpdateComment(comment *entities.Comment) (*entities.Comment, []error)
	DeleteComment(id uint) (*entities.Comment, []error)
	StoreComment(comment *entities.Comment) (*entities.Comment, []error)
}
