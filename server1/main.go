package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	db := database{"42": "life, universe and everything"}
	log.Fatal(http.ListenAndServe("localhost:8000", &db))
}

type database map[string]string

func (db *database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for k, v := range *db {
			fmt.Fprintf(w, "%s: %s", k, v)
		}
	case "/single":
		key := req.URL.Query().Get("key")
		if key == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Key not specified\n")
			return
		}
		v, ok := (*db)[key]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No such key: %s\n", key)
			return
		}
		fmt.Fprintf(w, "%s\n", v)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such page: %s\n", req.URL)
	}
}
