package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(workerId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Println("Worker ID: ", workerId)
		response, error := http.Get(site.URL)
		if error != nil {
			log.Println(error)
		}

		results <- Result{
			URL:    site.URL,
			Status: response.StatusCode,
		}
	}
}

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for worker := 1; worker <= 3; worker++ {
		go crawl(worker, jobs, results)
	}

	urls := []string{
		"https://apple.com",
		"https://google.com",
		"https://twitter.com",
		"https://netflix.com",
	}

	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for i := 1; i <= 4; i++ {
		result := results
		log.Println(result)
	}
}
