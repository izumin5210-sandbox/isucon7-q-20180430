package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var (
	fileRoot   string
	fileServer http.Handler
)

func main() {
	os.Exit(run())
}

func run() int {
	fileRoot = os.Getenv("ISUBATA_FILE_ROOT")
	if fileRoot == "" {
		fileRoot = "data"
	}
	fileServer = http.FileServer(http.Dir(fileRoot))
	err := http.ListenAndServe(":8080", http.HandlerFunc(handle))
	if err != nil {
		log.Printf("failed to shutdown server: %v", err)
		return 1
	}
	return 0
}

func handle(w http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s", req.Method, req.URL.Path)
	switch req.Method {
	case "GET":
		w.Header().Set("Vary", "Accept-Encoding")
		w.Header().Set("Cache-Control", "public, max-age=31557600")
		fileServer.ServeHTTP(w, req)
	case "POST":
		req.ParseMultipartForm(32 << 40)
		src, _, err := req.FormFile("image")
		if err != nil {
			log.Printf("failed to open form: %v\n", err)
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		defer src.Close()
		dst, err := os.Create(filepath.Join(fileRoot, req.URL.Path))
		if err != nil {
			log.Printf("failed to create file: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer dst.Close()
		_, err = io.Copy(dst, src)
		if err != nil {
			log.Printf("failed to copy file: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}
}
