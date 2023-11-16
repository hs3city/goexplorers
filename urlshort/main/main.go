package main

import (
	"fmt"
	urlshort "goexplorers/urlshort/main"
	"log/slog"
	"net/http"
	"os"
)

const urlMapYaml = "urlshort/urls.yaml"
const urlMapJson = "urlshort/urls.json"

var logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yamlFile, err := os.ReadFile(urlMapYaml)
	if err != nil {
		panic(err)
	}

	jsonFile, err := os.ReadFile(urlMapJson)
	if err != nil {
		panic(err)
	}

	yaml := string(yamlFile)
	json := string(jsonFile)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	logger.Info("Starting the server on :8080")
	http.ListenAndServe(":8080", jsonHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}
