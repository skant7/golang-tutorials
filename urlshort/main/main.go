package main

import (
	"fmt"
	"net/http"
	"urlshort"
)

//construct a defaultMux where fallback is a simple hello world
//chain MapHandler and yamlHandler in the mux
//default fallback is hello world
func main() {
	mux := defaultMux()
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)

	if err != nil {
		panic(err)
	}
	fmt.Println("Starting Server on Port http://localhost:8080")
	http.ListenAndServe(":8080", yamlHandler)
}

//yaml handler falls back to map Handler which fall backs to deafult

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

//deafult fallback
func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello from default mux")
}
