package routes

import (
	controllers "github.com/guancang10/BookStore/API/controllers/category"
	"github.com/julienschmidt/httprouter"
)

func SetRouter(controller controllers.CategoryController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", controller.GetAll)
	router.GET("/api/categories/:category_id", controller.Get)
	router.POST("/api/categories", controller.Save)
	router.DELETE("/api/categories/:category_id", controller.Delete)
	router.PUT("/api/categories/:category_id", controller.Update)

	return router
}
