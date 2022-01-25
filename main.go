package main

import (
	"github.com/Fachrurazzi/golang-restful-api/app"
	"github.com/Fachrurazzi/golang-restful-api/controller"
	"github.com/Fachrurazzi/golang-restful-api/helper"
	"github.com/Fachrurazzi/golang-restful-api/middleware"
	"github.com/Fachrurazzi/golang-restful-api/repository"
	"github.com/Fachrurazzi/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3333",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
