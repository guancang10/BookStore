package middleware

import (
	"github.com/guancang10/BookStore/API/helper/converter"
	response "github.com/guancang10/BookStore/API/models/web/response"
	"net/http"
)

type Middleware struct {
	handler http.Handler
}

func NewMiddleware(handler http.Handler) *Middleware {
	return &Middleware{handler: handler}
}

func (m Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	if request.Header.Get("x-api-key") == "Token" {
		m.handler.ServeHTTP(writer, request)
	} else {
		writer.Header().Add("content-type", "application/json")
		writer.WriteHeader(http.StatusUnauthorized)
		res := response.ApiResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}
		converter.EncoderToResponse(writer, res)
	}
}
