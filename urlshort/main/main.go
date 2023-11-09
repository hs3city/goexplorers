package main

import (
	"fmt"
	urlshort "goexplorers/urlshort/main"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	mux := defaultMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux, logger)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlFile, errReader := os.ReadFile("urlshort/urls.yaml")
	if errReader != nil {
		panic(errReader)
	}

	yaml := string(yamlFile)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler, logger)
	if err != nil {
		panic(err)
	}
	logger.Info("Starting the server on :8080")
	// http.ListenAndServe(":8080", mapHandler)
	http.ListenAndServe(":8080", yamlHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
