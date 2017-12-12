package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zacscodingclub/go-your-own-way"
)

func main() {
	filename := flag.String("file", "data.json", "the JSON file with the Go Your Own Way story")
	flag.Parse()

	f, err := os.Open(*filename)
	_ = f
	errorExit(err)

	story, err := gyow.JsonStory(f)
	errorExit(err)

	fmt.Printf("%+v\n", story)
}

func errorExit(e error) {
	if e != nil {
		fmt.Printf("Could not %v\n", e)
		os.Exit(1)
	}
}
