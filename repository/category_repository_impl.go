package repository

import (
	"context"
	"database/sql"
	"errors"
	"golang_restful_api/helper"
	"golang_restful_api/model/domain"
)

type CategoryRepositoryImpl struct {
}

// best practice nya yaitu membuat provider function tetap return ke struct nya bukan ke interface nya
func NewCategoryRepositoryImpl() *CategoryRepositoryImpl {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "INSERT INTO category(name, create_time, update_time) VALUES(?, ?, ?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, category.CreateTime, category.UpdateTime)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	category.Id = int(id)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ?, update_time = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.UpdateTime, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	SQL := "SELECT id, name, create_time, update_time FROM category WHERE id = ? LIMIT 1"
	rows, err := tx.QueryContext(ctx, SQL, categoryId)
	helper.PanicIfError(err)
	defer rows.Close()

	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name, &category.CreateTime, &category.UpdateTime)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name, create_time, update_time FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.CreateTime, &category.UpdateTime)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}

	return categories
}
