package main

import (
	"fmt"
	"html/template"
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
//
// First write simple text to the response.
//
// Optional if you have time: try html.template for the /list handler.
// * template.New("foo").Parse(`<html>...</html>`)
// * template.Execute
// * Try injecting html code through db values.

func handleDefault(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "No such page: %s\n", req.URL)
}

const layout = `
<html><body>
{{range $key, $value := .}} <li> {{$key}} <br/> {{$value}} </li> {{end}}
</body></html>`

func list(w http.ResponseWriter, req *http.Request) {
	// for i, j := range db {
	// 	fmt.Fprintf(w, "%s: %s\n", i, j)
	// }
	template, err := template.New("foo").Parse(layout)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	template.Execute(w, db)
}

func single(w http.ResponseWriter, req *http.Request) {
	key := req.URL.Query().Get("key")
	if key == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Key not specified, add &key= to the url\n")
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
