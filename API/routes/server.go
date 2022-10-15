package routes

import (
	"github.com/guancang10/BookStore/API/middleware"
	"net/http"
)

func SetServer(middleware *middleware.Middleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:8080",
		Handler: middleware,
	}
}
