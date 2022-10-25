package routes

import (
	controllers "github.com/guancang10/BookStore/API/controllers"
	"github.com/guancang10/BookStore/API/helper/exception"
	"github.com/julienschmidt/httprouter"
)

// Task : Nanti cari tau ini gimana optimaisenya, boros banget masa kalo ada controller lebih dari 1 di define 1 1
func SetRouter(category controllers.CategoryController, book controllers.BookController, user controllers.UserController, transaction controllers.TransactionController) *httprouter.Router {
	router := httprouter.New()
	router.GET("/api/categories", category.GetAll)
	router.GET("/api/categories/:category_id", category.Get)
	router.POST("/api/categories", category.Save)
	router.DELETE("/api/categories/:category_id", category.Delete)
	router.PUT("/api/categories/:category_id", category.Update)

	router.GET("/api/books", book.GetAll)
	router.GET("/api/books/:book_id", book.Get)
	router.POST("/api/books", book.Save)
	router.PUT("/api/books/:book_id", book.Update)
	router.PUT("/api/books/:book_id/add", book.AddQuantity)
	router.PUT("/api/books/:book_id/sub", book.SubQuantity)
	router.DELETE("/api/books/:book_id", book.Delete)

	router.GET("/api/users/:username", user.GetUser)
	router.GET("/api/users", user.GetAllUser)
	router.POST("/api/users/register", user.Register)
	router.POST("/api/users/login", user.Login)
	router.POST("/api/users/update", user.UpdateProfile)
	router.POST("/api/users/changepassword", user.ChangePassword)

	router.GET("/api/transactions/getheader/:username", transaction.GetTransactionHeaderUser)
	router.GET("/api/transactions/getdetail/:htr_book_id", transaction.GetTransactionHeaderDetail)
	router.POST("/api/transactions/create", transaction.CreateTransaction)
	router.POST("/api/transactions/update/status", transaction.UpdateTransactionStatus)
	router.POST("/api/transactions/update", transaction.UpdateTransaction)

	router.PanicHandler = exception.ErrorHandler

	return router
}
