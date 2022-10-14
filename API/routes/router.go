package routes

import (
	controllers "github.com/guancang10/BookStore/API/controllers/category"
	"github.com/julienschmidt/httprouter"
)

func SetRouter(controller controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", controller.GetAll)
}
