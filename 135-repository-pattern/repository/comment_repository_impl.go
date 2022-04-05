package repository

import (
	"135-repository-pattern/entity"
	"context"
	"database/sql"
	"errors"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	/*
		kenapa kembaliannya &commentRepositoryImpl{DB: db} ???
		karena method Insert, FindAll dan FindById adalah method pointer of struct commentRepositoryImpl,
		kalo commentRepositoryImpl bukan pointer return commentRepositoryImpl biasa juga bisa
	*/
	return &commentRepositoryImpl{DB: db}
}

func (repo *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	sqlQuery := "INSERT INTO comments(email, comment) VALUES (?, ?)"
	result, err := repo.DB.ExecContext(ctx, sqlQuery, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}

	comment.Id = int32(id)
	return comment, nil
}

func (repo *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	var comment entity.Comment
	sqlQuery := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repo.DB.QueryContext(ctx, sqlQuery, id)
	if err != nil {
		return comment, nil
	}
	defer rows.Close()

	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		return comment, errors.New("not found id " + strconv.Itoa(int(id)))
	}
}

func (repo *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	var comments []entity.Comment

	sqlQuery := "SELECT id, email, comment FROM comments"
	rows, err := repo.DB.QueryContext(ctx, sqlQuery)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var comment entity.Comment

		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}

	return comments, nil
}
