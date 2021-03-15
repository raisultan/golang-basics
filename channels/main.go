package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls := []string{
		"http://google.com",
		"http://fb.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.org",
	}

	// create channel with the type of string
	c := make(chan string)

	for _, url := range urls {
		// birth of new child - goroutine
		go checkUrl(url, c)
	}

	// receiving data from channel and processing it
	for url := range c {
		go func(u string) { // function literal = anon func
			time.Sleep(3 * time.Second)
			checkUrl(u, c)
		}(url) // pass by value not reference, because original value can be changed during the processing of the func
	}
	// equivalent for [just syntactic sugar for simplicity of comprehension]
	// continuous for loop
	// for {
	// 	go checkUrl(<-c, c)
	// }
}

// channel must be passed to func to communicate back and must be fully typed - `c chan string`
func checkUrl(url string, c chan string) {
	_, err := http.Get(url) // blocking call
	if err != nil {
		fmt.Println("Something wrong with - " + url)
	} else {
		fmt.Println("Everything is OK with - " + url)
	}
	c <- url
}

// everything that executed in golang is executed in goroutines (main func too)
// when the compiler sees keyword `go` before function call, it creates distinct goroutine for that function
// so it takes func content and executes it in distinct micro-process, in our case, it makes request to specific url
// waits until http.Get() returns smth and when it returns, which means execution is finished, the resulting value of func
// ran in distinct goroutine returns to main goroutine and will be handled there
// overall compileflow of main func is not changed. though, when it sees `go` keyword, compiler just executes the func
// and moves on, which results in number of goroutines at the same time.
// main difference is results of those executed funcs returns to main, only when corresponding goroutine finishes its execution

// low level explanation goroutines
// by default golang compiler executes goroutines only on one core of CPU
// this process handled by Go Scheduler. its purpose is to monitor and detect when to run specific goroutine, when to wait
// and execute another goroutine, while one is on the blocking process.

// with multiple CPU cores enabled for golang compiler Go Scheduler starts to work a bit different.
// the difference is that now Go Scheduler can run multiple goroutines at the same time, based on number of cores.
// for example, if machine has 3 enabled cores, then scheduler runs 3 goroutines at once.
// contrary for one core, when scheduler runs one at a time, and waits for blocking process to launch another goroutine.

// CONCURRENCY[thread switches on the fly] & PARALLELLISM[multiple things at the same time]

// Concurrency means that we can have multiple threads on which we switch execution, based on thread's state.
// in other words, if one thread starts blocking process, we switch to another thread. concisely, it means that we have smart thread scheduler.

// Parallellism is about processing and doing multiple things - goroutines at the same time, which clearly states that we must have MULTIPLE cores.

// Goroutines priority & Channels
// main func always has first level of priority, that means, it controls the workflow of the whole program, when to start and when to exit.
// so there can be a subtle misunderstanding. for example, in our case, when main func finishes looping through the slice, it finishes.
// because there is nothing else to run. even though we have multiple goroutines, those are not finished their execution.
// so here comes Channels to help us COMMUNICATE BETWEEN GOROUTINES. and Channels is the only way to communicate between goroutines.
// Channel is value, but its behavior and working mechanism is different. it is pub/sub like mechanism.
// if goroutine has access to a channel, it can receive values passed to the channel by other goroutines.
// and we should not forget that channels are strictly typed. so data passed and received through the channel can only be of one specifc type.

// note: receiving values from channel is blocking process, so we have to wait for values to receive from channel.
