package controllers

import (
	"context"
	"github.com/guancang10/BookStore/API/helper/converter"
	"github.com/guancang10/BookStore/API/models/web/request"
	response "github.com/guancang10/BookStore/API/models/web/response"
	"github.com/guancang10/BookStore/API/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService services.UserServices
}

func NewUserControllerImpl(service services.UserServices) UserController {
	return &UserControllerImpl{service}
}

func (u UserControllerImpl) Register(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var userRequest request.UserRegisterRequest
	converter.DecoderFromRequest(req, &userRequest)

	u.UserService.Register(context.Background(), userRequest)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success register account",
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (u UserControllerImpl) Login(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var userLoginRequest request.UserLoginRequest
	converter.DecoderFromRequest(req, &userLoginRequest)

	result := u.UserService.CheckPassword(context.Background(), userLoginRequest)
	//cookies := http.Cookie{
	//	Name:    "user-login",
	//	Value:   result.Username,
	//	Expires: time.Now().Add(5 * time.Minute),
	//}
	//http.SetCookie(writer, &cookies)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success login",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (u UserControllerImpl) Logout(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	// api not provide the logout, logout will provided in web apps
}

func (u UserControllerImpl) GetUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	username := params.ByName("username")
	result := u.UserService.GetUser(context.Background(), username)

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success get user detail",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (u UserControllerImpl) GetAllUser(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	result := u.UserService.GetAllUser(context.Background())

	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success get user",
		Data:   result,
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (u UserControllerImpl) ChangePassword(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var userChangePass request.UserChangePassword
	converter.DecoderFromRequest(req, &userChangePass)

	u.UserService.UpdatePassword(context.Background(), userChangePass)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success change password",
	}
	converter.EncoderToResponse(writer, webResponse)
}

func (u UserControllerImpl) UpdateProfile(writer http.ResponseWriter, req *http.Request, params httprouter.Params) {
	var updateRequest request.UserUpdateRequest
	converter.DecoderFromRequest(req, &updateRequest)

	u.UserService.UpdateProfile(context.Background(), updateRequest)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success update profile",
	}
	converter.EncoderToResponse(writer, webResponse)
}
