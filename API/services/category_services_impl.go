package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/helper/exception"
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/repository"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	Db                 *sql.DB
	Validator          *validator.Validate
}

func NewCategoryServiceImpl(categoryRepository repository.CategoryRepository, db *sql.DB, validator *validator.Validate) CategoryServices {
	return &CategoryServiceImpl{categoryRepository, db, validator}
}

func (c CategoryServiceImpl) Save(ctx context.Context, request request.CategoryCreateRequest) response.CategoryCreateResponse {
	fmt.Print(request)
	err := c.Validator.Struct(request)
	helper.CheckError(err)
	tx, err := c.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)

	category := domain.Category{
		CategoryName:  request.CategoryName,
		AuditUsername: request.AuditUsername,
	}
	category = c.CategoryRepository.Save(ctx, tx, category)

	return converter.ToCategoryCreateResponse(category)
}

func (c CategoryServiceImpl) GetAll(ctx context.Context) []response.CategoryGetResponse {
	tx, err := c.Db.Begin()
	helper.CheckError(err)

	categories := c.CategoryRepository.GetAll(ctx, tx)
	var response []response.CategoryGetResponse
	for _, v := range categories {
		response = append(response, converter.ToCategoryGetResponse(v))
	}
	return response
}

func (c CategoryServiceImpl) Get(ctx context.Context, categoryId int) response.CategoryGetResponse {
	tx, err := c.Db.Begin()
	helper.CheckError(err)

	category, err := c.CategoryRepository.Get(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}
	return converter.ToCategoryGetResponse(category)
}

func (c CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.Db.Begin()
	helper.CheckError(err)
	defer helper.CheckErrorTx(tx)
	_, err = c.CategoryRepository.Get(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	c.CategoryRepository.Delete(ctx, tx, categoryId)
}

func (c CategoryServiceImpl) Update(ctx context.Context, categoryId int, request request.CategoryUpdateRequest) {
	err := c.Validator.Struct(request)
	helper.CheckError(err)
	tx, err := c.Db.Begin()
	helper.CheckError(err)

	defer helper.CheckErrorTx(tx)
	_, err = c.CategoryRepository.Get(ctx, tx, categoryId)
	if err != nil {
		panic(exception.NewNotFoundError(err))
	}

	category := domain.Category{
		CategoryName:  request.CategoryName,
		AuditUsername: request.AuditUsername,
	}
	c.CategoryRepository.Update(ctx, tx, categoryId, category)
}
