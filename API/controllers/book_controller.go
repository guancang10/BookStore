package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type BookController interface {
	Save(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	AddQuantity(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	SubQuantity(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Get(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetAll(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}
