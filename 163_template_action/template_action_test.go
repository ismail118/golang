package _63_template_action

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.html"))

	t.ExecuteTemplate(w, "if.html", map[string]interface{}{
		"Name": "Eko",
	})
}

func TestTemplate(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateComparation(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparation.html"))

	t.ExecuteTemplate(w, "comparation.html", map[string]interface{}{
		"FinalValue": 80,
	})
}

func TestTemplateComparation(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateComparation(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.html"))

	t.ExecuteTemplate(w, "range.html", map[string]interface{}{
		"Hobbies": []string{
			"tidur", "makan", "game",
		},
	})
}

func TestTemplateRange(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateRange(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TemplateWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.html"))

	t.ExecuteTemplate(w, "address.html", map[string]interface{}{
		"Name": "Ismail",
		"Address": map[string]interface{}{
			"Street": "jl nin aja",
			"City":   "depok",
		},
	})
}

func TestTemplateWith(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	TemplateWith(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
