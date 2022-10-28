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

type PaymentControllerImpl struct {
	PaymentService services.PaymentService
}

func NewPaymentControllerImpl(paymentService services.PaymentService) PaymentController {
	return &PaymentControllerImpl{PaymentService: paymentService}
}

func (p PaymentControllerImpl) CreatePayment(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var paymentCreateRequest request.PaymentCreateRequest
	converter.DecoderFromRequest(req, &paymentCreateRequest)

	result := p.PaymentService.CreatePayment(context.Background(), paymentCreateRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Create Payment",
		Data:   result,
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (p PaymentControllerImpl) GetPaymentDetail(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	sPaymentId := params.ByName("payment_id")
	paymentId, err := strconv.Atoi(sPaymentId)
	helper.CheckError(err)

	result := p.PaymentService.GetPaymentDetail(context.Background(), paymentId)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Get Payment Detail",
		Data:   result,
	}

	converter.EncoderToResponse(writer, webResponse)

}

func (p PaymentControllerImpl) UpdatePaymentType(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var paymentUpdateRequest request.PaymentUpdateTypeRequest
	converter.DecoderFromRequest(req, &paymentUpdateRequest)

	p.PaymentService.UpdatePaymentType(context.Background(), paymentUpdateRequest)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Update Payment Type",
	}

	converter.EncoderToResponse(writer, webResponse)
}

func (p PaymentControllerImpl) UpdatePaymentStatus(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var paymentUpdateRequest request.PaymentUpdateStatusRequest
	converter.DecoderFromRequest(req, &paymentUpdateRequest)

	p.PaymentService.UpdatePaymentStatus(context.Background(), paymentUpdateRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success Update Payment Status",
	}

	converter.EncoderToResponse(writer, webResponse)
}
