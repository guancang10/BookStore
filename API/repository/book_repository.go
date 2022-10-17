package repository

import (
	"context"
	"database/sql"
	"github.com/guancang10/BookStore/API/models/domain"
)

type BookRepository interface {
	Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book
	AddQuantity(ctx context.Context, tx *sql.Tx, bookId int, qty int)
	SubQuantity(ctx context.Context, tx *sql.Tx, bookId int, qty int)
	Get(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Book
	Update(ctx context.Context, tx *sql.Tx, bookId int, book domain.Book)
	Delete(ctx context.Context, tx *sql.Tx, bookId int)
}
