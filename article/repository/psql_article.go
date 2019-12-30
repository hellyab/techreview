package repository

import (
	"database/sql"
	"errors"

	"github.com/hellyab/techreview/entity"
)

type ArticleRepositoryImpl struct { // contains sql pointer
	conn *sql.DB
}

func NewArticleRepositoryImpl(Conn *sql.DB) *ArticleRepositoryImpl { //
	return &ArticleRepositoryImpl{conn: Conn}
}

// AricleRepositoryImpl implements aricle.ArticleRepository

func (ari *ArticleRepositoryImpl) Articles() ([]entity.Article, error) {
	rows, err := ari.conn.Query("SELECT * FROM articles;")

	if err != nil {
		return nil, errors.New("cannot query the database")
	}
	defer rows.Close()

	articles := []entity.Article{}

	for rows.Next() {
		article := entity.Article{}

		err = rows.Scan(&article.ID, &article.Name)

		if err != nil {
			return nil, err
		}

		articles = append(articles, article)

	}

	return articles, nil
}
