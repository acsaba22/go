package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const defaultid = "L7vObOAUAPoC"
const url_base = "https://www.googleapis.com/books/v1/volumes/"

// Google books has a REST API, let's use it! Example link:
// https://www.googleapis.com/books/v1/volumes/L7vObOAUAPoC
// Description of the API:
// https://developers.google.com/books/docs/v1/reference/volumes#resource
//
// res, err := http.Get(url) gives you a http.Response
//
// ioutil.ReadAll will take res.Body and return []byte
//
// json.Unmarshal will take []byte and a struct varible
// and fill the matching field names from the json.
// Get these fields:
//   id, (for helping)
//   volumeInfo.title,
//   volumeinfo.authors,
//   volumeinfo.categories,
//
// Reminder:
// type book struct {
//	  fieldname fieldtype
// }
// b := book{}
//
// Print out the title, author, category
//
// If you still have time make the id an argument:
// Get os.Args[1] (if exists), and use that instead of the default id.
//
// Other ids to check:
// vrPQAwAAQBAJ EVmminvaWDQC 5oSU5PepogEC

type book struct {
	Id         string
	VolumeInfo struct {
		Title      string
		Authors    []string
		Categories []string
		Spoiler    string `json:"description"`
	}
}

func main() {
	id := defaultid
	if 2 <= len(os.Args) {
		id = os.Args[1]
	}
	url := url_base + id
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while fetching: %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error while reading Body: %v", err)
		return
	}

	// fmt.Println(string(data), err)
	var b book
	err = json.Unmarshal(data, &b)
	if err != nil {
		fmt.Println("Error while unmarshal: ", err)
	}
	// fmt.Printf("%+v\n", b)
	fmt.Printf("Title: %s\nAuthors: %s\nCategories: %s\n",
		b.VolumeInfo.Title,
		strings.Join(b.VolumeInfo.Authors, ", "),
		strings.Join(b.VolumeInfo.Categories, ", "))
}
