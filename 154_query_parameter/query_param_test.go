package _54_query_parameter

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func SayHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		fmt.Fprint(w, "Hello")
	} else {
		fmt.Fprintf(w, "Hello %s", name)
	}
}

func TestQueryParam(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/hello?name=eko", nil)
	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameter(w http.ResponseWriter, r *http.Request) {
	firstName := r.URL.Query().Get("first_name")
	lastName := r.URL.Query().Get("last_name")

	fmt.Fprintf(w, "Fist Name: %s \nLast Name: %s \n", firstName, lastName)
}

func TestMultipleParameter(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/hi?first_name=ismail&last_name=alfiyasin", nil)
	recorder := httptest.NewRecorder()

	MultipleParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

func MultipleParameterValues(w http.ResponseWriter, r *http.Request) {
	var query map[string][]string = r.URL.Query()
	var names []string = query["names"]
	fmt.Fprintln(w, strings.Join(names, ","))
}

func TestMultipleParameterValues(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost/hi-all?names=ismail&names=edo", nil)
	recorder := httptest.NewRecorder()

	MultipleParameterValues(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
