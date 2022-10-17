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

func InitServer() *http.Server {
	wire.Build(
		appdb.GetConnection,
		validator.New,
		repository.NewCategoryRepository,
		services.NewCategoryServiceImpl,
		controllers.NewCategoryControllerImpl,
		repository.NewBookRepositoryImpl,
		services.NewBookServiceImpl,
		controllers.NewBookControllerImpl,
		routes.SetRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewMiddleware,
		routes.SetServer,
	)
	return nil
}
