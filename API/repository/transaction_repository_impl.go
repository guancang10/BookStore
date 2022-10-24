package repository

import (
	"context"
	"database/sql"
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

func (t TransactionRepositoryImpl) SaveTransaction(ctx context.Context, tx *sql.Tx, trBook domain.TrBook) {
	script := "INSERT INTO trbook(HtrBookId,BookId,Price,Qty,AuditUsername) VALUES(?,?,?,?,?)"
	_, err := tx.ExecContext(ctx, script, trBook.HtrBookId, trBook.BookId, trBook.Price, trBook.Qty, trBook.AuditUsername)
	helper.CheckError(err)
}

//func (t TransactionRepositoryImpl) GetHeaderTr(ctx context.Context, tx *sql.Tx, htrBookId int) (domain.HtrBook, error) {
//	script := "SELECT Id,Username,TotalPrice,TransactionDate,AuditUsername FROM htrbook WHERE Id = ?"
//	row, err := tx.QueryContext(ctx, script, htrBookId)
//	helper.CheckError(err)
//	var result domain.HtrBook
//	defer row.Close()
//	if row.Next() {
//		err := row.Scan(result.Id, result.Username, result.TotalPrice, result.TransactionDate, result.AuditUsername)
//		helper.CheckError(err)
//		return result, nil
//	} else {
//		return result, errors.New("transaction not exists")
//	}
//}

//func (t TransactionRepositoryImpl) GetAllHeaderTr(ctx context.Context, tx *sql.Tx) []domain.HtrBook {
//	script := "SELECT Id,Username,TotalPrice,TransactionDate,AuditUsername FROM htrbook"
//	rows, err := tx.QueryContext(ctx, script)
//	helper.CheckError(err)
//
//	defer rows.Close()
//	var result []domain.HtrBook
//	for rows.Next() {
//		var data domain.HtrBook
//		err := rows.Scan(&data.Id, &data.Username, &data.TotalPrice, &data.TransactionDate, &data.AuditUsername)
//		helper.CheckError(err)
//		result = append(result, data)
//	}
//	return result
//}

func (t TransactionRepositoryImpl) GetHeaderTrUser(ctx context.Context, tx *sql.Tx, username string) []domain.HtrBook {
	script := "SELECT Id,TransactionDate,Username,TotalPrice FROM htrbook WHERE Username =?"
	rows, err := tx.QueryContext(ctx, script, username)
	helper.CheckError(err)

	defer rows.Close()
	var result []domain.HtrBook
	for rows.Next() {
		var data domain.HtrBook
		err := rows.Scan(&data.Id, &data.TransactionDate, &data.Username, &data.TotalPrice)
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
