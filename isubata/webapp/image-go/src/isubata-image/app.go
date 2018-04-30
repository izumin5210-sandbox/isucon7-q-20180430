package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	fileRoot string
)

func main() {
	os.Exit(run())
}

func run() int {
	fileRoot = os.Getenv("ISUCON_FILE_ROOT")
	if fileRoot == "" {
		fileRoot = "./data"
	}
	err := http.ListenAndServe(":8080", http.HandlerFunc(handle))
	if err != nil {
		log.Printf("failed to shutdown server: %v", err)
		return 1
	}
	return 0
}

func handle(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		http.ServeFile(w, req, filepath.Join(fileRoot, req.URL.Path))
	case "POST":
		src, _, err := req.FormFile("image")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer src.Close()
		dst, err := os.Create(filepath.Join(fileRoot, req.URL.Path))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
