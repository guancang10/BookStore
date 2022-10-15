//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/guancang10/BookStore/API/appdb"
	controllers "github.com/guancang10/BookStore/API/controllers/category"
	"github.com/guancang10/BookStore/API/middleware"
	repository "github.com/guancang10/BookStore/API/repository/category"
	"github.com/guancang10/BookStore/API/routes"
	services "github.com/guancang10/BookStore/API/services/category"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func InitServer() *http.Server {
	wire.Build(
		appdb.GetConnection,
		validator.New,
		repository.NewCategoryRepository,
		services.NewCategoryServiceImpl,
		controllers.NewCategoryController,
		routes.SetRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewMiddleware,
		routes.SetServer,
	)
	return nil
}
