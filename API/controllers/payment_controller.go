package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type PaymentController interface {
	CreatePayment(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetPaymentDetail(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	UpdatePaymentType(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	UpdatePaymentStatus(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}
