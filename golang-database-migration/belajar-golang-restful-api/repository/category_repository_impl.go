package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"kursus/belajar-golang-restful-api/helper"
	"kursus/belajar-golang-restful-api/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySql := "INSERT INTO category2(name) values (?)"
	result, err := tx.ExecContext(ctx, querySql, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	querySql := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (c *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, category domain.Category) {
	querySql := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, querySql, category.Id)
	helper.PanicIfError(err)
}

func (c *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	querySql := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, querySql, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	var category domain.Category

	if rows.Next() {
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New(fmt.Sprintf("category with id = %d not found", categoryId))
	}
}

func (c *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	querySql := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, querySql)
	helper.PanicIfError(err)
	defer rows.Close()

	var categorys []domain.Category

	for rows.Next() {
		var category domain.Category
		err = rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categorys = append(categorys, category)
	}

	return categorys
}
