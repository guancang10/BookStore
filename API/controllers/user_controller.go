package controllers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Register(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Login(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	Logout(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	GetAllUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	ChangePassword(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
	UpdateProfile(writer http.ResponseWriter, req *http.Request, params httprouter.Params)
}
