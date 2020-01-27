package service

import (
	"github.com/hellyab/techreview/comment"
	"github.com/hellyab/techreview/entities"
)

type CommentService struct {
	commentRepo comment.CommentRepository
}

func NewCommentService(ComRepo comment.CommentRepository) comment.CommentService {
	return &CommentService{commentRepo: ComRepo}
}

func (cs *CommentService) Comments() ([]entities.Comment, []error) {
	comments, errs := cs.commentRepo.Comments()
	if len(errs) > 0 {
		return nil, errs
	}

	return comments, nil
}

func (cs *CommentService) Comment(id uint) (*entities.Comment, []error) {
	cmnt, errs := cs.commentRepo.Comment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cs *CommentService) StoreComment(comment *entities.Comment) (*entities.Comment, []error) {
	cmnt, errs := cs.commentRepo.StoreComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cs *CommentService) UpdateComment(comment *entities.Comment) (*entities.Comment, []error) {
	cmnt, errs := cs.commentRepo.UpdateComment(comment)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cs *CommentService) DeleteComment(id uint) (*entities.Comment, []error) {
	cmnt, errs := cs.commentRepo.DeleteComment(id)
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
