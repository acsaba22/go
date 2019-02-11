package main

import (
	"fmt"
	"log"
	"net/http"
)

var db = map[string]string{
	"42":       "life, universe and everything",
	"question": "What do you get if you multiply six by nine?",
}

func main() {
	http.HandleFunc("/", handleDefault)
	http.HandleFunc("/list", list)
	http.HandleFunc("/single", single)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Given a w of type http.ResponseWriter
// this is how you write the result page:
// fmt.Fprintf(w, "%s", s)
//
// You can also set headers to w:
// w.WriteHeader(http.StatusNotFound) or
// w.WriteHeader(http.StatusBadRequest)
//
// Given req of type http.Request
// you can req.URL.Query().Get("key")

func handleDefault(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "No such page: %s\n", req.URL)
}

func list(w http.ResponseWriter, req *http.Request) {
	for k, v := range db {
		fmt.Fprintf(w, "%s: %s\n", k, v)
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
	fmt.Fprintf(w, "%s\n", v)
}
