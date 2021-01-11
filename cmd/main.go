package main

import (
	"fmt"
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
}
