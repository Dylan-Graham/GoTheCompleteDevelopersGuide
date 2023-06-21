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

	c := make(chan string)

	for _, url := range links {
		go checkStatus(url, c)
		// fmt.Println(<-c)
	}

	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkStatus(url string, c chan string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("%s might be down\n", url)
		c <- "Might be down I think"
	}

	if response.StatusCode == 200 {
		fmt.Printf("%s is okay\n", url)
		c <- "Site is okay"
	}
}

// Go-Routines: go functionThatYouWantToRunInParellel
// Go-Channels: Allow go routines to communicate

// Create a channel: c := make(chan string)

// Send into channel: channel <- 5
// Wait for value in channel: myNumber <- channel
// Wait for value in channel & log: fmt.Println(<- channel)
