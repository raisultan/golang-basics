package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://fb.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.org",
	}

	for _, url := range urls {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		if resp.StatusCode == 200 {
			fmt.Println("Everything is OK with - ", url)
		} else {
			fmt.Println("Something wrong with - ", url)
		}
	}

}
