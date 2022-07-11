package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string {
	"https://twitter.com",
	"https://google.com",
	"https://apple.com",
}

func fetch(url string, waitGroup *sync.WaitGroup) {
	response, error := http.Get(url)
	if error != nil {
		println(error)
	}
	fmt.Println(response.Status)
	waitGroup.Done()
}

func crawl() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(1)
	for _, url := range urls {
		waitGroup.Add(1)
		go fetch(url, &waitGroup)
	}
	waitGroup.Wait()
	println("Finished crawling")
}

func main() {
	crawl()
}