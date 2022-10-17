package exception

import (
	"compress/flate"
	"github.com/go-playground/validator/v10"
	"github.com/guancang10/BookStore/API/helper/converter"
	response2 "github.com/guancang10/BookStore/API/models/web/response"
	"net/http"
)

// NotFoundError struct for make a new type of error
type NotFoundError struct {
	err string
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{err: err.Error()}
}

func ErrorHandler(writer http.ResponseWriter, req *http.Request, err interface{}) {
	if BadRequest(writer, req, err) {
		return
	} else if NotFound(writer, req, err) {
		return
	} else if InternalServerError(writer, req, err) {
		return
	} else if UnknownError(writer, req, err) {
		return
	}
}

func BadRequest(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, fit := err.(validator.ValidationErrors)
	if fit {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		response := response2.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		converter.EncoderToResponse(writer, response)
		return true
	} else {
		return false
	}
}

func NotFound(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, fit := err.(NotFoundError)
	if fit {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusNotFound)
		response := response2.ApiResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   exception.err,
		}
		converter.EncoderToResponse(writer, response)
		return true
	} else {
		return false
	}
}

func InternalServerError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, fit := err.(flate.InternalError)
	if fit {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusInternalServerError)
		webResponse := response2.ApiResponse{
			Code:   http.StatusInternalServerError,
			Status: "INTERNAL SERVER ERROR",
			Data:   exception.Error(),
		}
		converter.EncoderToResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}

func UnknownError(writer http.ResponseWriter, req *http.Request, err interface{}) bool {
	exception, fit := err.(error)
	if fit {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)
		webResponse := response2.ApiResponse{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   exception.Error(),
		}
		converter.EncoderToResponse(writer, webResponse)
		return true
	} else {
		return false
	}
}
