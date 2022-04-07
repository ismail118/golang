package _57_response_code

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "name is empty")
	} else {
		// tidak perlu di set status code nya jika 200 karena sudah defaultnya 200
		fmt.Fprintf(w, "Hi %s", name)
	}
}

func TestResponseCodeValid(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080?name=eko", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}

func TestResponseCodeFailed(t *testing.T) {
	request := httptest.NewRequest(http.MethodPost, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	ResponseCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
}
