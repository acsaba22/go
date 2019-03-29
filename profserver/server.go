package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/count", handleCount)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handleCount(w http.ResponseWriter, req *http.Request) {
	keyStr := req.URL.Query().Get("n")
	if keyStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "n not specified, try ?n=5\n")
		return
	}
	key, err := strconv.Atoi(keyStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "n not an integer: %d", key)
		return
	}
	ret := 0
	for i := 0; i < key; i++ {
		ret++
	}
	fmt.Fprintf(w, "%v", ret)
}
