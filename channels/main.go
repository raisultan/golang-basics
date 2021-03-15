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
	_, err := http.Get(url) // blocking call
	if err != nil {
		fmt.Println("Something wrong with - ", url)
		return
	}

	fmt.Println("Everything is OK with - ", url)
}

// everything that executed in golang is executed in goroutines (main func too)
// when the compiler sees keyword `go` before function call, it creates distinct goroutine for that function
// so it takes func content and executes it in distinct micro-process, in our case, it makes request to specific url
// waits until http.Get() returns smth and when it returns, which means execution is finished, the resulting value of func
// ran in distinct goroutine returns to main goroutine and will be handled there
// overall compileflow of main func is not changed. though, when it sees `go` keyword, compiler just executes the func
// and moves on, which results in number of goroutines at the same time.
// main difference is results of those executed funcs returns to main, only when corresponding goroutine finishes its execution
