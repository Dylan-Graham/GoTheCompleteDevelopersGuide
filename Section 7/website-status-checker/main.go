package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
	}

	for _, url := range links {
		checkStatus(url)
	}
}

func checkStatus(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s might be down\n", url)
	}

	if response.StatusCode == 200 {
		fmt.Printf("%s is okay\n", url)
	}
}
