package services

import (
	"context"
	"database/sql"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
)

type TransactionServiceImpl struct {
	Db                    *sql.DB
	Validator             *validator.Validate
	TransactionRepository repository.TransactionRepository
	BookRepository repository.BookRepository
}

func NewTransactionServiceImpl(db *sql.DB, validator *validator.Validate, transactionrepository repository.TransactionRepository, bookRepository repository.BookRepository) TransactionService {
	return &TransactionServiceImpl{Db: db, Validator: validator, TransactionRepository: transactionrepository, BookRepository: bookRepository}
}

func (t TransactionServiceImpl) CreateTransaction(ctx context.Context, transaction request.TransactionInsertRequest) response.TransactionInsertResponse {
	err := t.Validator.Struct(transaction)
	helper.CheckError(err)

	tx, err := t.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)

	transactionDetailReq := transaction.Detail
	var totalPrice float64
	for _, v := range transactionDetailReq {
		totalPrice += v.Price * float64(v.Qty)
	}

	htrBook := converter.FromCreateReqToHtrBook(transaction, totalPrice)
	htrBook.StatusId = 1
	htrBook = t.TransactionRepository.CreateHeaderTr(ctx, tx, htrBook)

	for _, v := range transactionDetailReq {
		trBook := converter.FromCreateReqToTrBook(v, htrBook.Id)
		t.TransactionRepository.SaveTr	ansaction(ctx, tx, trBook)
	}

}

func (t TransactionServiceImpl) GetTransactionHeaderUser(ctx context.Context, username string) []response.TransactionGetHeaderResponse {
	//TODO implement me
	panic("implement me")
}

func (t TransactionServiceImpl) GetTransactionHeaderDetail(ctx context.Context, htrBookId int) response.TransactionDetailResponse {
	//TODO implement me
	panic("implement me")
}
