package main

import (
	"fmt"
	"github.com/go-playground/validator"
	"kursus/belajar-golang-restful-api/app"
	"kursus/belajar-golang-restful-api/controller"
	"kursus/belajar-golang-restful-api/helper"
	middleware "kursus/belajar-golang-restful-api/middleware"
	"kursus/belajar-golang-restful-api/repository"
	"kursus/belajar-golang-restful-api/services"
	"net/http"
)

func main() {

	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := services.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)

	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println(fmt.Sprintf("App run at %s", server.Addr))
}
