package main

const defaultID = "L7vObOAUAPoC"
const urlBase = "https://www.googleapis.com/books/v1/volumes/"

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

func main() {
}
