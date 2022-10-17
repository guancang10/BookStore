package routes

import (
	controllers "github.com/guancang10/BookStore/API/controllers"
	"github.com/guancang10/BookStore/API/helper/exception"
	"github.com/julienschmidt/httprouter"
)

// Task : Nanti cari tau ini gimana optimaisenya, boros banget masa kalo ada controller lebih dari 1 di define 1 1
func SetRouter(controller controllers.CategoryController, controller2 controllers.BookController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", controller.GetAll)
	router.GET("/api/categories/:category_id", controller.Get)
	router.POST("/api/categories", controller.Save)
	router.DELETE("/api/categories/:category_id", controller.Delete)
	router.PUT("/api/categories/:category_id", controller.Update)

	router.GET("/api/books", controller2.GetAll)
	router.GET("/api/books/:book_id", controller2.Get)
	router.POST("/api/books", controller2.Save)
	router.PUT("/api/books/:book_id", controller2.Update)
	router.PUT("/api/books/:book_id/add", controller2.AddQuantity)
	router.PUT("/api/books/:book_id/sub", controller2.SubQuantity)
	router.DELETE("/api/books/:book_id", controller2.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}
