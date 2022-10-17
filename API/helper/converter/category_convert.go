package converter

import (
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/response"
)

func ToCategoryCreateResponse(category domain.Category) response.CategoryCreateResponse {
	return response.CategoryCreateResponse{
		Id: category.Id, CategoryName: category.CategoryName, AuditUsername: category.AuditUsername,
	}
}

func ToCategoryGetResponse(category domain.Category) response.CategoryGetResponse {
	return response.CategoryGetResponse{
		Id: category.Id, CategoryName: category.CategoryName, AuditUsername: category.AuditUsername,
	}
}
