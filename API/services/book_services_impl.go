package services

import (
	"context"
	"database/sql"
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/helper/exception"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
	"net/http"
	"strconv"
)

type BookServicesImpl struct {
	Db             *sql.DB
	BookRepository repository.BookRepository
	Validator      *validator.Validate
}

func NewBookServiceImpl(db *sql.DB, bookRepository repository.BookRepository, validator *validator.Validate) BookServices {
	return &BookServicesImpl{Db: db, BookRepository: bookRepository, Validator: validator}
}

func (b BookServicesImpl) Save(ctx context.Context, book request.BookCreateRequest) response.BookCreateResponse {
	err := b.Validator.Struct(book)
	helper.CheckError(err)
	tx, err := b.Db.Begin()
	helper.CheckError(err)

	//Check category exists or not
	client := &http.Client{}
	url := "http://localhost:8080/api/categories/" + strconv.Itoa(book.CategoryId)
	req, err := http.NewRequest("GET", url, nil)
	helper.CheckError(err)
	req.Header.Add("x-api-key", "Token")
	res, err := client.Do(req)
	helper.CheckError(err)
	if res.StatusCode != 200 {
		panic(exception.NewNotFoundError(errors.New("category not exists")))
	}

	defer helper.CheckErrorTx(tx)
	param := converter.FromRequestToBook(book)
	result := b.BookRepository.Save(ctx, tx, param)

	return converter.FromBookToCreateResponse(result)
}

func (b BookServicesImpl) AddQuantity(ctx context.Context, bookId int, qty request.BookQuantityRequest) response.BookQuantityResponse {
	err := b.Validator.Struct(qty)
	helper.CheckError(err)
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	exists, err := b.BookRepository.Get(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	} else {
		b.BookRepository.AddQuantity(ctx, tx, bookId, qty.Qty)
		qtyResult := exists.Qty + qty.Qty
		return converter.FromQtytoQtyResponse(qtyResult)
	}
}

func (b BookServicesImpl) SubQuantity(ctx context.Context, bookId int, qty request.BookQuantityRequest) response.BookQuantityResponse {
	err := b.Validator.Struct(qty)
	helper.CheckError(err)
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	exists, err := b.BookRepository.Get(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	} else {
		qtyResult := exists.Qty - qty.Qty
		if qtyResult < 0 {
			panic(errors.New("can't subtract more than existing quantity"))
		} else {
			b.BookRepository.SubQuantity(ctx, tx, bookId, qty.Qty)
			return converter.FromQtytoQtyResponse(qtyResult)
		}
	}
}

func (b BookServicesImpl) Get(ctx context.Context, bookId int) response.BookCreateResponse {
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	result, err := b.BookRepository.Get(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	} else {
		return converter.FromBookToCreateResponse(result)
	}
}

func (b BookServicesImpl) GetAll(ctx context.Context) []response.BookCreateResponse {
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	result := b.BookRepository.GetAll(ctx, tx)
	return converter.FromSliceBookToCreateResponse(result)
}

func (b BookServicesImpl) Update(ctx context.Context, bookId int, book request.BookUpdateRequest) {
	err := b.Validator.Struct(book)
	helper.CheckError(err)
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)
	_, err = b.BookRepository.Get(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	} else {
		b.BookRepository.Update(ctx, tx, bookId, converter.FromRequestToBookUpdate(book))
	}
}

func (b BookServicesImpl) Delete(ctx context.Context, bookId int) {
	tx, err := b.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	_, err = b.BookRepository.Get(ctx, tx, bookId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	} else {
		b.BookRepository.Delete(ctx, tx, bookId)
	}
}
