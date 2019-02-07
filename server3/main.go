package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("../templates/*"))

func main() {
	fmt.Println("Names: ")
	for _, t := range templates.Templates() {
		fmt.Println(t.Name())
	}
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/list", list)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/single", single)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type database map[string]string

var db = database{"42": "life, universe and everything"}

func hello(w http.ResponseWriter, req *http.Request) {
	// err := templates.ExecuteTemplate(w, "hello.html", struct{}{})
	err := templates.ExecuteTemplate(w, "display1.html", struct {
		Key   string
		Value string
	}{"a", "b"})
	if err != nil {
		log.Fatal("Error during serving: ", err)
	}
}

func handleDefault(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "No such page: %s\n", req.URL)
}

func list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s", k, v)
	}
}

func single(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Key not specified\n")
		return
	}
	v, ok := db[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such key: %s\n", key)
		return
	}
	err := templates.ExecuteTemplate(w, "display1.html", struct {
		Key   string
		Value string
	}{key, v})
	if err != nil {
		log.Fatal("Error during serving: ", err)
	}
	// fmt.Fprintf(w, "%s\n", v)
}
