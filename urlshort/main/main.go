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

	yamlFile := catch(os.ReadFile(urlMapYaml))
	jsonFile := catch(os.ReadFile(urlMapJson))

	yaml := string(yamlFile)
	json := string(jsonFile)

	yamlHandler := catch(urlshort.YAMLHandler([]byte(yaml), mapHandler))
	jsonHandler := catch(urlshort.JSONHandler([]byte(json), yamlHandler))

	logger.Info("Starting the server on :8080")
	if err := http.ListenAndServe(":8080", jsonHandler); err != nil {
		logger.Error("Error starting server", err)
	}

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func catch[T any](val T, err error) T {
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	return val
}
