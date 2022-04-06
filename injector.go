// go:build wireinject
//go:build wireinject
// +build wireinject

package main

import (
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/middleware"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

var categorySet = wire.NewSet(
	// bind berarti jika ada yg membutuhkan interface categoryrepository, maka akan dikirimkan categoryrepositoryimpl
	// karena udah di-bind maka dia jadi bersatu
	repository.NewCategoryRepositoryImpl,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryServiceImpl,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryControllerImpl,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(
			new(http.Handler), new(*httprouter.Router),
		),
		middleware.NewAuthMiddleware,
		app.NewServer,
	)
	return nil
}
