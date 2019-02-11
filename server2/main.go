package main

import (
	"log"
	"net/http"
)

var db = map[string]string{
	"42":       "life, universe and everything",
	"question": "What do you get if you multiply six by nine?",
}

func main() {
	// http.HandleFunc("/", handleDefault)
	// http.HandleFunc("/list", list)
	// http.HandleFunc("/single", single)
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
