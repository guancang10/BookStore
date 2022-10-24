package repository

import (
	"context"
	"database/sql"
	"github.com/guancang10/BookStore/API/models/domain"
)

type TransactionRepository interface {
	CreateHeaderTr(ctx context.Context, tx *sql.Tx, htrBook domain.HtrBook) domain.HtrBook
	SaveTransaction(ctx context.Context, tx *sql.Tx, trBook domain.TrBook) domain.TrBook
	UpdateTransactionStatus(ctx context.Context, tx *sql.Tx, htrBookId int, statusId int, auditUsername string)
	GetHeaderTr(ctx context.Context, tx *sql.Tx, htrBookId int) (domain.HtrBook, error)
	GetHeaderTrUser(ctx context.Context, tx *sql.Tx, username string) []domain.HtrBook
	GetHeaderDetail(ctx context.Context, tx *sql.Tx, htrBookId int) []domain.TrBook
}
