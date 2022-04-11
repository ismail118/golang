package _70_download_file

import (
	"fmt"
	"net/http"
	"path"
	"testing"
)

func DownloadFile(w http.ResponseWriter, r *http.Request) {
	fileName := r.URL.Query().Get("file")
	if fileName == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "BAD REQUEST")
		return
	}

	w.Header().Add("Content-Disposition", "attachment; filename=\""+fileName+"\"")
	http.ServeFile(w, r, path.Join("./resources", fileName))
}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	server.ListenAndServe()
}
