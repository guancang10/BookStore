package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/models/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepositoryImpl() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (t TransactionRepositoryImpl) CreateHeaderTr(ctx context.Context, tx *sql.Tx, htrBook domain.HtrBook) domain.HtrBook {
	script := "INSERT INTO htrbook(Username,TotalPrice,TransactionDate,StatusId,AuditUsername) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, htrBook.Username, htrBook.TotalPrice, htrBook.TransactionDate, htrBook.StatusId, htrBook.AuditUsername)
	helper.CheckError(err)
	id, err := result.LastInsertId()
	helper.CheckError(err)
	htrBook.Id = int(id)

	return htrBook
}

func (t TransactionRepositoryImpl) SaveTransaction(ctx context.Context, tx *sql.Tx, trBook domain.TrBook) domain.TrBook {
	script := "INSERT INTO trbook(HtrBookId,BookId,Price,Qty,AuditUsername) VALUES(?,?,?,?,?)"
	result, err := tx.ExecContext(ctx, script, trBook.HtrBookId, trBook.BookId, trBook.Price, trBook.Qty, trBook.AuditUsername)
	id, err := result.LastInsertId()
	helper.CheckError(err)
	trBook.Id = int(id)

	helper.CheckError(err)
	return trBook
}

func (t TransactionRepositoryImpl) UpdateTransactionDetail(ctx context.Context, tx *sql.Tx, trBook domain.TrBook) {
	script := "UPDATE TrBook SET Qty = ?,AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, trBook.Qty, trBook.AuditUsername, trBook.Id)
	helper.CheckError(err)
}

func (t TransactionRepositoryImpl) UpdateTransactionHeader(ctx context.Context, tx *sql.Tx, htrBook domain.HtrBook) {
	script := "UPDATE HtrBook SET TotalPrice = ? ,AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, htrBook.TotalPrice, htrBook.AuditUsername, htrBook.Id)
	helper.CheckError(err)
}

func (t TransactionRepositoryImpl) UpdateTransactionStatus(ctx context.Context, tx *sql.Tx, htrBookId int, statusId int, auditUsername string) {
	script := "UPDATE htrbook SET StatusId = ?,AuditUsername = ? WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, statusId, auditUsername, htrBookId)
	helper.CheckError(err)
}

func (t TransactionRepositoryImpl) DeleteTransactionDetail(ctx context.Context, tx *sql.Tx, trBookId int) {
	script := "DELETE FROM TrBook WHERE Id = ?"
	_, err := tx.ExecContext(ctx, script, trBookId)
	helper.CheckError(err)
}

func (t TransactionRepositoryImpl) GetHeaderTr(ctx context.Context, tx *sql.Tx, htrBookId int) (domain.HtrBook, error) {
	script := "SELECT Id,Username,TotalPrice,TransactionDate,StatusId,AuditUsername FROM htrbook WHERE Id = ?"
	row, err := tx.QueryContext(ctx, script, htrBookId)
	helper.CheckError(err)
	var result domain.HtrBook
	defer row.Close()
	if row.Next() {
		err := row.Scan(&result.Id, &result.Username, &result.TotalPrice, &result.TransactionDate, &result.StatusId, &result.AuditUsername)
		helper.CheckError(err)
		return result, nil
	} else {
		return result, errors.New("transaction not exists")
	}
}

func (t TransactionRepositoryImpl) GetHeaderTrUser(ctx context.Context, tx *sql.Tx, username string) []domain.HtrBook {
	script := "SELECT Id,TransactionDate,Username,StatusId,TotalPrice,AuditUsername FROM htrbook WHERE Username =?"
	rows, err := tx.QueryContext(ctx, script, username)
	helper.CheckError(err)

	defer rows.Close()
	var result []domain.HtrBook
	for rows.Next() {
		var data domain.HtrBook
		err := rows.Scan(&data.Id, &data.TransactionDate, &data.Username, &data.StatusId, &data.TotalPrice, &data.AuditUsername)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}

func (t TransactionRepositoryImpl) GetHeaderDetail(ctx context.Context, tx *sql.Tx, htrBookId int) []domain.TrBook {
	script := "SELECT Id,HtrBookId,BookId,Price,Qty,AuditUsername FROM trbook WHERE HtrBookId = ?"
	rows, err := tx.QueryContext(ctx, script, htrBookId)
	helper.CheckError(err)

	defer rows.Close()
	var result []domain.TrBook
	for rows.Next() {
		var data domain.TrBook
		err := rows.Scan(&data.Id, &data.HtrBookId, &data.BookId, &data.Price, &data.Qty, &data.AuditUsername)
		helper.CheckError(err)
		result = append(result, data)
	}
	return result
}
