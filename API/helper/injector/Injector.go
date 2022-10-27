//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/guancang10/BookStore/API/appdb"
	"github.com/guancang10/BookStore/API/controllers"
	"github.com/guancang10/BookStore/API/middleware"
	"github.com/guancang10/BookStore/API/repository"
	"github.com/guancang10/BookStore/API/routes"
	"github.com/guancang10/BookStore/API/services"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var categoryController = wire.NewSet(
	repository.NewCategoryRepository,
	services.NewCategoryServiceImpl,
	controllers.NewCategoryControllerImpl,
)

var bookController = wire.NewSet(
	repository.NewBookRepositoryImpl,
	services.NewBookServiceImpl,
	controllers.NewBookControllerImpl,
)

var userController = wire.NewSet(
	repository.NewUserRepositoryImpl,
	services.NewUserServiceImpl,
	controllers.NewUserControllerImpl,
)

var transactionController = wire.NewSet(
	repository.NewTransactionRepositoryImpl,
	services.NewTransactionServiceImpl,
	controllers.NewTransactionControllerImpl,
)
var transactionController = wire.NewSet(
	repository.NewTransactionRepositoryImpl,
	services.NewTransactionServiceImpl,
	controllers.NewTransactionControllerImpl,
)
var transactionController = wire.NewSet(
	repository.NewTransactionRepositoryImpl,
	services.NewTransactionServiceImpl,
	controllers.NewTransactionControllerImpl,
)

func InitServer() *http.Server {
	wire.Build(
		appdb.GetConnection,
		validator.New,
		categoryController,
		bookController,
		userController,
		transactionController,
		routes.SetRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewMiddleware,
		routes.SetServer,
	)
	return nil
}
