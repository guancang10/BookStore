package repository

import (
	"context"
	"database/sql"
	"github.com/guancang10/BookStore/API/models/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Category
	Get(ctx context.Context, tx *sql.Tx, categoryId int) (domain.Category, error)
	Delete(ctx context.Context, tx *sql.Tx, categoryId int)
	Update(ctx context.Context, tx *sql.Tx, categoryId int, category domain.Category)
}
