package repository

import (
	"github.com/hellyab/techreview/comment"
	"github.com/hellyab/techreview/entities"
	"github.com/jinzhu/gorm"
)

type CommentGormRepo struct {
	conn *gorm.DB
}

func NewCommentGormRepo(db *gorm.DB) comment.CommentRepository {

	return &CommentGormRepo{conn: db}

}

func (cRepo *CommentGormRepo) Comments() ([]entities.Comment, []error) {
	comments := []entities.Comment{}
	errs := cRepo.conn.Find(&comments).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return comments, errs
}

func (cmntRepo *CommentGormRepo) StoreComment(comment *entities.Comment) (*entities.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Create(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cmntRepo *CommentGormRepo) Comment(id uint) (*entities.Comment, []error) {
	cmnt := entities.Comment{}
	errs := cmntRepo.conn.First(&cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return &cmnt, errs
}

func (cmntRepo *CommentGormRepo) UpdateComment(comment *entities.Comment) (*entities.Comment, []error) {
	cmnt := comment
	errs := cmntRepo.conn.Save(cmnt).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}

func (cmntRepo *CommentGormRepo) DeleteComment(id uint) (*entities.Comment, []error) {
	cmnt, errs := cmntRepo.Comment(id)

	if len(errs) > 0 {
		return nil, errs
	}

	errs = cmntRepo.conn.Delete(cmnt, id).GetErrors()
	if len(errs) > 0 {
		return nil, errs
	}
	return cmnt, errs
}
