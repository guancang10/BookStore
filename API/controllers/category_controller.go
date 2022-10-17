package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type CategoryController interface {
	//Parameter position must same (responsewritter,request,params) so we can insert it as handle in approtuer
	Save(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetAll(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Get(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Delete(res http.ResponseWriter, req *http.Request, params httprouter.Params)
	Update(res http.ResponseWriter, req *http.Request, params httprouter.Params)
}
