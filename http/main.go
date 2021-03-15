package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Number of bytes written: ", len(bs))

	return len(bs), nil
}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	lw := logWriter{}

	// reader interface writes to writer interface
	io.Copy(lw, resp.Body)
}
