//go:build wireinject
// +build wireinject

package main

import (
	"github.com/go-playground/validator"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
	"kursus/belajar-golang-dependency-injection/app"
	"kursus/belajar-golang-dependency-injection/controller"
	"kursus/belajar-golang-dependency-injection/middleware"
	"kursus/belajar-golang-dependency-injection/repository"
	"kursus/belajar-golang-dependency-injection/services"
	"net/http"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	services.NewCategoryService,
	wire.Bind(new(services.CategoryService), new(*services.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

func InitializedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)

	return nil
}
