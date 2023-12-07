package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func main() {
	fmt.Println("Server starting...")
	http.HandleFunc("/", handlerHttp)
	http.ListenAndServe(":8080", nil)
}

func handlerHttp(w http.ResponseWriter, r *http.Request) {

	pathSegments := strings.Split(r.URL.Path, "/")
	expectedArc := pathSegments[1]

	jsonInput, err := os.ReadFile("data.json")
	if err != nil {
		panic(err)
	}

	var story StoryLine

	err = json.Unmarshal(jsonInput, &story)
	if err != nil {
		panic(err)
	}

	t, _ := template.ParseFiles("template.html")

	val := reflect.ValueOf(story)
	typ := reflect.TypeOf(story)
	var retVal string

	for i := 0; i < val.NumField(); i++ {
		if typ.Field(i).Name == expectedArc {
			fmt.Println(val.Field(i))
			retVal = val.Field(i).String()
		}
	}

	err = t.Execute(w, retVal)

	if err != nil {
		panic(err)
	}
}
