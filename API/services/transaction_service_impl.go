package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
	"strconv"
)

type TransactionServiceImpl struct {
	Db                    *sql.DB
	Validator             *validator.Validate
	TransactionRepository repository.TransactionRepository
	BookRepository        repository.BookRepository
	UserRepository        repository.UserRepository
}

func NewTransactionServiceImpl(db *sql.DB, validator *validator.Validate, transactionRepository repository.TransactionRepository, bookRepository repository.BookRepository, userRepository repository.UserRepository) TransactionService {
	return &TransactionServiceImpl{Db: db, Validator: validator, TransactionRepository: transactionRepository, BookRepository: bookRepository, UserRepository: userRepository}
}

func (t TransactionServiceImpl) CreateTransaction(ctx context.Context, transaction request.TransactionInsertRequest) response.TransactionDetailResponse {
	err := t.Validator.Struct(transaction)
	helper.CheckError(err)

	tx, err := t.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)

	transactionDetailReq := transaction.Detail

	//check user
	_, err = t.UserRepository.GetUser(ctx, tx, transaction.Username)
	helper.CheckError(err)

	//Getting price and total price
	var totalPrice float64
	var mapBook = make(map[int]float64)
	for _, v := range transactionDetailReq {
		bookDetail, err := t.BookRepository.Get(ctx, tx, v.BookId)
		helper.CheckError(err)
		if bookDetail.Qty-v.Qty < 0 {
			stringError := "Can't subtract book with Id " + strconv.Itoa(v.BookId) + " more then available quantity"
			panic(errors.New(stringError))
		}

		mapBook[v.BookId] = bookDetail.Price
		totalPrice += bookDetail.Price * float64(v.Qty)
	}

	//Create header transaction
	htrBook := converter.FromCreateReqToHtrBook(transaction, totalPrice)
	htrBook.StatusId = 1
	htrBook = t.TransactionRepository.CreateHeaderTr(ctx, tx, htrBook)

	//save transaction
	var listTrBook []domain.TrBook
	for _, v := range transactionDetailReq {
		trBook := converter.FromCreateReqToTrBook(v, htrBook.Id, mapBook[v.BookId])
		trBook = t.TransactionRepository.SaveTransaction(ctx, tx, trBook)
		listTrBook = append(listTrBook, trBook)
	}

	return converter.CreateTransactionDetailResponse(htrBook, listTrBook)
}

func (t TransactionServiceImpl) UpdateTransaction(ctx context.Context, transaction request.TransactionUpdateRequest) response.TransactionDetailResponse {
	err := t.Validator.Struct(transaction)
	helper.CheckError(err)

	tx, err := t.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)

	trBookReq := transaction.Detail

	//Update transaction detail
	for _, v := range trBookReq {
		trBook := converter.FromUpdateReqToTrBook(v)
		if trBook.Qty == 0 {
			t.TransactionRepository.DeleteTransactionDetail(ctx, tx, trBook.Id)
		} else {
			t.TransactionRepository.UpdateTransactionDetail(ctx, tx, trBook)
		}
	}

	//calculate new total price
	trBookList := t.TransactionRepository.GetHeaderDetail(ctx, tx, transaction.HtrBookId)
	var totalPrice float64
	for _, v := range trBookList {
		bookDetail, err := t.BookRepository.Get(ctx, tx, v.BookId)
		helper.CheckError(err)
		if v.Qty < 0 || v.Qty > bookDetail.Qty {
			stringError := "can't set quantity below 0 or above available quantity"
			panic(errors.New(stringError))
		}
		totalPrice += float64(v.Qty) * v.Price
	}

	//update header transaction
	htrBook, err := t.TransactionRepository.GetHeaderTr(ctx, tx, transaction.HtrBookId)
	helper.CheckError(err)
	htrBook.TotalPrice = totalPrice
	htrBook.AuditUsername = transaction.AuditUsername
	t.TransactionRepository.UpdateTransactionHeader(ctx, tx, htrBook)

	return converter.CreateTransactionDetailResponse(htrBook, trBookList)
}

func (t TransactionServiceImpl) UpdateTransactionStatus(ctx context.Context, status request.TransactionUpdateStatusReq) {
	err := t.Validator.Struct(status)
	helper.CheckError(err)

	tx, err := t.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	htrBook, err := t.TransactionRepository.GetHeaderTr(ctx, tx, status.HtrBookId)
	helper.CheckError(err)

	if htrBook.StatusId == 3 || htrBook.StatusId == 4 {
		panic(errors.New("can't update success or canceled transaction"))
	}

	if status.StatusId == 3 {
		trBookList := t.TransactionRepository.GetHeaderDetail(ctx, tx, status.HtrBookId)
		for _, v := range trBookList {
			bookDetail, err := t.BookRepository.Get(ctx, tx, v.BookId)
			helper.CheckError(err)
			if bookDetail.Qty-v.Qty < 0 {
				stringError := "Can't subtract book with Id " + strconv.Itoa(v.BookId) + " more then available quantity"
				panic(errors.New(stringError))
			}
			t.BookRepository.SubQuantity(ctx, tx, v.BookId, v.Qty)
		}
	}

	t.TransactionRepository.UpdateTransactionStatus(ctx, tx, status.HtrBookId, status.StatusId, status.AuditUsername)
}

func (t TransactionServiceImpl) GetTransactionHeaderUser(ctx context.Context, username string) []response.TransactionGetHeaderResponse {
	tx, err := t.Db.Begin()
	helper.CheckError(err)

	//check user
	_, err = t.UserRepository.GetUser(ctx, tx, username)
	helper.CheckError(err)

	result := t.TransactionRepository.GetHeaderTrUser(ctx, tx, username)
	return converter.FromArrHeaderBookToResponse(result)
}

func (t TransactionServiceImpl) GetTransactionHeaderDetail(ctx context.Context, htrBookId int) response.TransactionDetailResponse {
	tx, err := t.Db.Begin()
	helper.CheckError(err)

	htrBook, err := t.TransactionRepository.GetHeaderTr(ctx, tx, htrBookId)
	helper.CheckError(err)

	trBook := t.TransactionRepository.GetHeaderDetail(ctx, tx, htrBookId)

	return converter.CreateTransactionDetailResponse(htrBook, trBook)
}
