package controllers

import (
	"context"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/web/request"
	"github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BookControllerImpl struct {
	BookService services.BookServices
}

func NewBookControllerImpl(bookService services.BookServices) BookController {
	return &BookControllerImpl{BookService: bookService}
}

func (b BookControllerImpl) Save(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var bookRequest request.BookCreateRequest
	converter.DecoderFromRequest(req, &bookRequest)
	bookResponse := b.BookService.Save(context.Background(), bookRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Insert New Book",
		Data:   bookResponse,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) AddQuantity(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var bookRequest request.BookQuantityRequest
	converter.DecoderFromRequest(req, &bookRequest)
	bookId, err := strconv.Atoi(params.ByName("book_id"))
	helper.CheckError(err)
	bookResponse := b.BookService.AddQuantity(context.Background(), bookId, bookRequest)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success add quantity",
		Data:   "New Quantity: " + strconv.Itoa(bookResponse.Qty),
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) SubQuantity(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var bookRequest request.BookQuantityRequest
	converter.DecoderFromRequest(req, &bookRequest)
	bookId, err := strconv.Atoi(params.ByName("book_id"))
	helper.CheckError(err)
	bookResponse := b.BookService.SubQuantity(context.Background(), bookId, bookRequest)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success subtract quantity",
		Data:   "New Quantity: " + strconv.Itoa(bookResponse.Qty),
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) Get(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	bookId, err := strconv.Atoi(params.ByName("book_id"))
	helper.CheckError(err)
	bookResponse := b.BookService.Get(context.Background(), bookId)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success get book",
		Data:   bookResponse,
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) GetAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	bookResponse := b.BookService.GetAll(context.Background())

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success get book",
		Data:   bookResponse,
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var bookRequest request.BookUpdateRequest
	converter.DecoderFromRequest(req, &bookRequest)
	bookId, err := strconv.Atoi(params.ByName("book_id"))
	helper.CheckError(err)
	b.BookService.Update(context.Background(), bookId, bookRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success update book",
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (b BookControllerImpl) Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	bookId, err := strconv.Atoi(params.ByName("book_id"))
	helper.CheckError(err)
	b.BookService.Delete(context.Background(), bookId)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success delete book",
	}

	converter.EncoderToResponse(writer, webResponse)
}
