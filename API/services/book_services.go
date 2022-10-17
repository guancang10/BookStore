package services

import (
	"context"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

type BookServices interface {
	Save(ctx context.Context, book request.BookCreateRequest) response.BookCreateResponse
	AddQuantity(ctx context.Context, bookId int, qty request.BookQuantityRequest) response.BookQuantityResponse
	SubQuantity(ctx context.Context, bookId int, qty request.BookQuantityRequest) response.BookQuantityResponse
	Get(ctx context.Context, bookId int) response.BookCreateResponse
	GetAll(ctx context.Context) []response.BookCreateResponse
	Update(ctx context.Context, bookId int, book request.BookUpdateRequest)
	Delete(ctx context.Context, bookId int)
}
