package main

import (
	"fmt"
	"kursus/belajar-golang-dependency-injection/helper"
	middleware "kursus/belajar-golang-dependency-injection/middleware"
	"net/http"
)

func NewServer(authMidlleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMidlleware,
	}
}
func main() {
	server := InitializedServer()

	err := server.ListenAndServe()
	helper.PanicIfError(err)

	fmt.Println(fmt.Sprintf("App run at %s", server.Addr))
}
