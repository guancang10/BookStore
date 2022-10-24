package services

import (
	"context"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

type TransactionService interface {
	CreateTransaction(ctx context.Context, transaction request.TransactionInsertRequest) response.TransactionInsertResponse
	GetTransactionHeaderUser(ctx context.Context, username string) []response.TransactionGetHeaderResponse
	GetTransactionHeaderDetail(ctx context.Context, htrBookId int) response.TransactionDetailResponse
}
