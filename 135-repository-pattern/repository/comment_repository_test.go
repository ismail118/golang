package repository

import (
	"135-repository-pattern"
	"135-repository-pattern/entity"
	"context"
	"fmt"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	db := _35_repository_pattern.GetConnection()
	defer db.Close()

	repository := NewCommentRepository(db)
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "ismail@gmail.com",
		Comment: "test comment",
	}

	result, err := repository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	db := _35_repository_pattern.GetConnection()
	defer db.Close()

	repository := NewCommentRepository(db)
	ctx := context.Background()
	id := int32(1)

	comment, err := repository.FindById(ctx, id)
	if err != nil {
		panic(err)
	}

	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	db := _35_repository_pattern.GetConnection()
	defer db.Close()

	repository := NewCommentRepository(db)
	ctx := context.Background()

	comments, err := repository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, each := range comments {
		fmt.Println(each)
	}
}
