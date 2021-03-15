package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		log.Fatal(err)
	}

	// reader interface writes to writer interface
	io.Copy(os.Stdout, resp.Body)
}
