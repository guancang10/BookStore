package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
)

type BookRepositoryImpl struct {
}

func NewBookRepositoryImpl() BookRepository {
	return &BookRepositoryImpl{}
}

func (b BookRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, book domain.Book) domain.Book {
	script := "INSERT INTO Book(BookName,BookDescription,Author,CategoryId,Qty,AuditUsername) VALUES(?,?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, book.BookName, book.BookDescription, book.Author, book.CategoryId, book.Qty, book.AuditUsername)
	helper.CheckError(err)
	id, err := result.LastInsertId()
	helper.CheckError(err)
	book.Id = int(id)
	return book
}

func (b BookRepositoryImpl) AddQuantity(ctx context.Context, tx *sql.Tx, bookId int, qty int) {
	script := "UPDATE Book SET Qty = Qty + ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, qty, bookId)
	helper.CheckError(err)

}

func (b BookRepositoryImpl) SubQuantity(ctx context.Context, tx *sql.Tx, bookId int, qty int) {
	script := "UPDATE Book SET Qty = Qty - ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, qty, bookId)
	helper.CheckError(err)
}

func (b BookRepositoryImpl) Get(ctx context.Context, tx *sql.Tx, bookId int) (domain.Book, error) {
	script := "SELECT Id,BookName,BookDescription,Author,CategoryId,Qty FROM Book WHERE Id = ?"
	rows, err := tx.QueryContext(ctx, script, bookId)
	helper.CheckError(err)
	defer rows.Close()
	var result domain.Book
	if rows.Next() {
		err := rows.Scan(&result.Id, &result.BookName, &result.BookDescription, &result.Author, &result.CategoryId, &result.Qty)
		helper.CheckError(err)
		return result, nil
	} else {
		return result, errors.New("book not found")
	}
}

func (b BookRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Book {
	script := "SELECT Id,BookName,BookDescription,Author,CategoryId,Qty From Book"
	rows, err := tx.QueryContext(ctx, script)
	helper.CheckError(err)
	defer rows.Close()
	var result []domain.Book
	for rows.Next() {
		var data domain.Book
		rows.Scan(&data.Id, &data.BookName, &data.BookDescription, &data.Author, &data.CategoryId, &data.Qty)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}

func (b BookRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, bookId int, book domain.Book) {
	script := "UPDATE Book SET BookName = ?,BookDescription = ?,Author = ?, CategoryId = ? , Qty = ?, AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, book.BookName, book.BookDescription, book.Author, book.CategoryId, book.Qty, book.AuditUsername)
	helper.CheckError(err)
}

func (b BookRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, bookId int) {
	script := "DELETE FROM Book WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, bookId)
	helper.CheckError(err)
}
