package main

import (
	"fmt"
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
		checkUrl(url)
	}

}

func checkUrl(url string) {
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Something wrong with - ", url)
		return
	}

	fmt.Println("Everything is OK with - ", url)
}
