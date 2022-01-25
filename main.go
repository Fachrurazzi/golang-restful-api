package main

import (
	"github.com/Fachrurazzi/golang-restful-api/app"
	"github.com/Fachrurazzi/golang-restful-api/controller"
	"github.com/Fachrurazzi/golang-restful-api/exception"
	"github.com/Fachrurazzi/golang-restful-api/helper"
	"github.com/Fachrurazzi/golang-restful-api/middleware"
	"github.com/Fachrurazzi/golang-restful-api/repository"
	"github.com/Fachrurazzi/golang-restful-api/service"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := httprouter.New()

	router.GET("/api/categories", categoryController.FindAll)
	router.GET("/api/categories/:categoryId", categoryController.FindById)
	router.POST("/api/categories", categoryController.Create)
	router.PUT("/api/categories/:categoryId", categoryController.Update)
	router.DELETE("/api/categories/:categoryId", categoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	server := http.Server{
		Addr:    "localhost:3333",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}