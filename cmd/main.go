package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/tkircsi/link"
)

func main() {
	f, err := os.Open("ex2.html")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	links, err := link.Parse(f)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", links)

	url := "http://golang.org"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	links, err = link.Parse(resp.Body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", links)
}
