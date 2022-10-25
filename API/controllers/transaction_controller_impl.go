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

type TransactionControllerImpl struct {
	TransactionService services.TransactionService
}

func NewTransactionControllerImpl(transactionService services.TransactionService) TransactionController {
	return &TransactionControllerImpl{TransactionService: transactionService}
}

func (t TransactionControllerImpl) CreateTransaction(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var transactionRequest request.TransactionInsertRequest
	converter.DecoderFromRequest(req, &transactionRequest)

	result := t.TransactionService.CreateTransaction(context.Background(), transactionRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Create Transaction",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (t TransactionControllerImpl) UpdateTransactionStatus(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var transactionUpdate request.TransactionUpdateStatusReq
	converter.DecoderFromRequest(req, &transactionUpdate)

	t.TransactionService.UpdateTransactionStatus(context.Background(), transactionUpdate)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Update Status",
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (t TransactionControllerImpl) UpdateTransaction(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var transactionUpdate request.TransactionUpdateRequest
	converter.DecoderFromRequest(req, &transactionUpdate)

	result := t.TransactionService.UpdateTransaction(context.Background(), transactionUpdate)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Update Transaction",
		Data:   result,
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (t TransactionControllerImpl) GetTransactionHeaderUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	username := params.ByName("username")

	result := t.TransactionService.GetTransactionHeaderUser(context.Background(), username)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Transaction Header",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (t TransactionControllerImpl) GetTransactionHeaderDetail(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	htrBookId, err := strconv.Atoi(params.ByName("htr_book_id"))
	helper.CheckError(err)

	result := t.TransactionService.GetTransactionHeaderDetail(context.Background(), htrBookId)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Transaction Detail",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}
