// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-playground/validator/v10"
	"golang_restful_api/app"
	"golang_restful_api/controller"
	"golang_restful_api/middleware"
	"golang_restful_api/repository"
	"golang_restful_api/service"
	"net/http"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from injector.go:

func InitServer() *http.Server {
	categoryRepositoryImpl := repository.NewCategoryRepositoryImpl()
	db := app.NewDB()
	validate := validator.New()
	categoryServiceImpl := service.NewCategoryServiceImpl(categoryRepositoryImpl, db, validate)
	categoryControllerImpl := controller.NewCategoryControllerImpl(categoryServiceImpl)
	router := app.NewRouter(categoryControllerImpl)
	authMiddleware := middleware.NewAuthMiddleware(router)
	server := app.NewServer(authMiddleware)
	return server
}
