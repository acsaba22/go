package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/acsaba22/go/iowordfreq"
)

const fileName = "urls.txt"

var wordsToKnow = strings.Split("html prize body script css", " ")

func getUrls() (urls []string, err error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urls = append(urls, scanner.Text())
	}
	err = scanner.Err()
	return
}

func main() {
	start := time.Now()
	urls, err := getUrls()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	wf := iowordfreq.IoWordFreq{}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, &wf, ch)
	}
	for range urls {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs Total\n", time.Since(start).Seconds())
	fmt.Println("==========================")
	for _, w := range wordsToKnow {
		fmt.Println(fmt.Sprintf("%s: %v", w, wf.GetWordCount(w)))
	}
}

func fetch(url string, w io.Writer, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	nbytes, err := io.Copy(w, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	t := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", t, nbytes, url)
}
