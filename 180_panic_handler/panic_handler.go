package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, i interface{}) {
		fmt.Fprint(writer, "Panic ", i)
	}

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		panic("Ups")
	})
}
