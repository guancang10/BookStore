package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type TransactionController interface {
	CreateTransaction(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	UpdateTransactionStatus(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetTransactionHeaderUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetTransactionHeaderDetail(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}