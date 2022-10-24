package converter

import (
	"github.com/guancang10/BookStore/API/models/domain"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
)

func FromRequestToBook(param request.BookCreateRequest) domain.Book {
	return domain.Book{
		BookName:        param.BookName,
		BookDescription: param.BookDescription,
		Author:          param.Author,
		Qty:             param.Qty,
		Price:           param.Price,
		CategoryId:      param.CategoryId,
		AuditUsername:   param.AuditUsername,
	}
}

func FromRequestToBookUpdate(param request.BookUpdateRequest) domain.Book {
	return domain.Book{
		BookName:        param.BookName,
		BookDescription: param.BookDescription,
		Author:          param.Author,
		Qty:             param.Qty,
		Price:           param.Price,
		CategoryId:      param.CategoryId,
		AuditUsername:   param.AuditUsername,
	}
}

func FromBookToCreateResponse(param domain.Book) response.BookCreateResponse {
	return response.BookCreateResponse{
		Id:              param.Id,
		BookName:        param.BookName,
		BookDescription: param.BookDescription,
		Author:          param.Author,
		CategoryId:      param.CategoryId,
		Qty:             param.Qty,
		Price:           param.Price,
	}
}

func FromQtytoQtyResponse(qty int) response.BookQuantityResponse {
	return response.BookQuantityResponse{
		Qty: qty,
	}
}

func FromSliceBookToCreateResponse(params []domain.Book) []response.BookCreateResponse {
	var result []response.BookCreateResponse
	for _, v := range params {
		result = append(result, FromBookToCreateResponse(v))
	}
	return result
}
