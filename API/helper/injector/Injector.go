//go:build wireinject
// +build wireinject

package injector

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/guancang10/BookStore/API/appdb"
	controllers "github.com/guancang10/BookStore/API/controllers/category"
	repository "github.com/guancang10/BookStore/API/repository/category"
	services "github.com/guancang10/BookStore/API/services/category"
	"net/http"
)

func InitServer() *http.Server {
	wire.Build(
		appdb.GetConnection(),
		validator.New(),
		repository.NewCategoryRepository(),
		services.NewCategoryServiceImpl(),
		controllers.NewCategoryController(),

	)
}
