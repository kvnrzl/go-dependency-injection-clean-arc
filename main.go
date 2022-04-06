package main

import (
	"golang_restful_api/helper"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	server := InitServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}

// func main() {
// 	// DB
// 	DB := app.NewDB()

// 	// validator
// 	validate := validator.New()

// 	// inject
// 	categoryRepository := repository.NewCategoryRepository()
// 	categoryService := service.NewCategoryService(categoryRepository, DB, validate)
// 	categoryController := controller.NewCategoryController(categoryService)

// 	// router
// 	router := app.NewRouter(categoryController)

// 	// server
// 	server := http.Server{
// 		Addr:    "localhost:3000",
// 		Handler: middleware.NewAuthMiddleware(router),
// 	}
// 	errServer := server.ListenAndServe()
// 	helper.PanicIfError(errServer)
// }
