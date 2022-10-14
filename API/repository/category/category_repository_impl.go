package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

// ExecContext for exec query without rows return(such as save delete,etc)
// QueryContext for exec query with rows return(get, getall,etc)
func (c CategoryRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	script := "INSERT INTO CATEGORY(CategoryName,AuditUsername) VALUES(?,?)"
	result, err := tx.ExecContext(ctx, script, category.CategoryName, category.AuditUsername)
	helper.CheckError(err)

	lastId, err := result.LastInsertId()
	helper.CheckError(err)

	category.Id = int(lastId)
	return category
}

func (c CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	query := "SELECT * FROM Category"
	rows, err := tx.QueryContext(ctx, query)
	defer rows.Close() // don't forget to close rows
	helper.CheckError(err)
	var result []domain.Category
	for rows.Next() {
		var data domain.Category
		err := rows.Scan(&data.Id, &data.CategoryName, &data.AuditUsername, &data.AuditTime)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}

func (c CategoryRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error) {
	script := "SELECT * FROM Category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, script, categoryId)
	helper.CheckError(err)
	var result domain.Category
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.CategoryName, &result.AuditUsername, &result.AuditTime)
		helper.CheckError(err)
		return result, nil
	} else {
		return result, errors.New("data not found")
	}

}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, categoryId int) {
	script := "DELETE Category WHERE id = ?"
	_, err := tx.ExecContext(ctx, script, categoryId)
	helper.CheckError(err)
}

func (c CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, categoryId int, category domain.Category) {
	script := "UPDATE Category SET CateogryName = ? , AuditUsername = ? WHERE id = ? "
	_, err := tx.ExecContext(ctx, script, category.CategoryName, category.AuditUsername, category.Id)
	helper.CheckError(err)
}
