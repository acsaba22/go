package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

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
	id := "L7vObOAUAPoC"
	if 2 <= len(os.Args) {
		id = os.Args[1]
	}
	url := "https://www.googleapis.com/books/v1/volumes/" + id
	// fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error while fetching: %v", err)
		return
	}
	data, err := ioutil.ReadAll(res.Body)

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
