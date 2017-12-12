package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zacscodingclub/go-your-own-way"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the GOYW web application")
	filename := flag.String("file", "data.json", "the JSON file with the Go Your Own Way story")
	flag.Parse()

	f, err := os.Open(*filename)
	_ = f
	errorAndExit(err)

	story, err := gyow.JsonStory(f)
	errorAndExit(err)

	h := gyow.NewHandler(story)
	fmt.Printf("Starting server at port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))

}

func errorAndExit(e error) {
	if e != nil {
		fmt.Printf("Could not %v\n", e)
		os.Exit(1)
	}
}
