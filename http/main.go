package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	response, err := http.Get("http://www.google.com")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}

	lw := logWriter{}
	io.Copy(lw, response.Body)
}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}
