package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"urlshortner"
)

func main() {
	yamlInput := flag.String("yaml", "./_redirects.yml", "YAML file containing the URL mapping")
	flag.Parse()

	yamlFile, err := ioutil.ReadFile(*yamlInput)
	if err != nil {
		log.Fatal("Error loading file "+*yamlInput, err)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshortner.MapHandler(pathsToUrls, mux)
	
	yamlHandler, err := urlshortner.YAMLHandler(yamlFile, mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
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
