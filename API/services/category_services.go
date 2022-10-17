package services

import (
	"context"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

type CategoryServices interface {
	Save(ctx context.Context, request request.CategoryCreateRequest) response.CategoryCreateResponse
	GetAll(ctx context.Context) []response.CategoryGetResponse
	Get(ctx context.Context, categoryId int) response.CategoryGetResponse
	Delete(ctx context.Context, categoryId int)
	Update(ctx context.Context, categoryId int, request request.CategoryUpdateRequest)
}
