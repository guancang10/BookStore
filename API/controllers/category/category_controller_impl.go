package controllers

import (
	"context"
	"github.com/guancang10/BookStore/API/helper"
	"github.com/guancang10/BookStore/API/helper/converter"
	request "github.com/guancang10/BookStore/API/models/web/request"
	response "github.com/guancang10/BookStore/API/models/web/response"
	services "github.com/guancang10/BookStore/API/services/category"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type CategoryControllerImpl struct {
	service services.CategoryServices
}

func NewCategoryController(service services.CategoryServices) CategoryController {
	return &CategoryControllerImpl{service: service}
}

func (c CategoryControllerImpl) Save(req *http.Request, res http.ResponseWriter, params httprouter.Params) {
	var categoryRequest request.CategoryCreateRequest
	converter.DecoderFromRequest(req, &categoryRequest)
	result := c.service.Save(context.Background(), categoryRequest)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   result,
	}
	converter.EncoderToResponse(res, webResponse)
}

func (c CategoryControllerImpl) GetAll(req *http.Request, res http.ResponseWriter, params httprouter.Params) {
	result := c.service.GetAll(context.Background())
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   result,
	}
	converter.EncoderToResponse(res, webResponse)
}

func (c CategoryControllerImpl) Get(req *http.Request, res http.ResponseWriter, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("category_id"))
	helper.CheckError(err)
	result := c.service.Get(context.Background(), categoryId)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
		Data:   result,
	}
	converter.EncoderToResponse(res, webResponse)
}

func (c CategoryControllerImpl) Delete(req *http.Request, res http.ResponseWriter, params httprouter.Params) {
	categoryId, err := strconv.Atoi(params.ByName("category_id"))
	helper.CheckError(err)
	c.service.Delete(context.Background(), categoryId)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
	}
	converter.EncoderToResponse(res, webResponse)
}

func (c CategoryControllerImpl) Update(req *http.Request, res http.ResponseWriter, params httprouter.Params) {
	var request request.CategoryUpdateRequest
	converter.DecoderFromRequest(req, request)
	categoryId, err := strconv.Atoi(params.ByName("category_id"))
	helper.CheckError(err)
	c.service.Update(context.Background(), categoryId, request)
	webResponse := response.ApiResponse{
		Code:   http.StatusOK,
		Status: "Success",
	}
	converter.EncoderToResponse(res, webResponse)
}
