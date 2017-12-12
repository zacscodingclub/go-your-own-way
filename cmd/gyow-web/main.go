package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

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
	tpl := template.Must(template.New("").Parse(storyTemplate))
	h := gyow.NewHandler(story, gyow.WithTemplate(tpl), gyow.WithPathFunc(pathFn))

	mux := http.NewServeMux()
	mux.Handle("/story/", h)

	fmt.Printf("Starting server at port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), mux))
}

func errorAndExit(e error) {
	if e != nil {
		fmt.Printf("Could not %v\n", e)
		os.Exit(1)
	}
}

func pathFn(r *http.Request) string {
	path := strings.TrimSpace(r.URL.Path)
	if path == "/story" || path == "/story/" {
		path = "/story/intro"
	}

	return path[len("/story/"):]
}

var storyTemplate = `
<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<meta http-equiv="X-UA-Compatible" content="ie=edge">
		<title>Go Your Own Way</title>
	</head>
	<body>
		<section class="page">
			<h1>{{.Title}}</h1>

			{{range .Paragraphs}}
				<p>{{.}}</p>
			{{end}}

			<ul>
			{{range .Options}}
				<li><a href="/story/{{.Chapter}}">{{.Text}}</a></li>
			{{end}}
			</ul>
		</section>
	</body>
	<style>
		body {
			font-family: helvetica, arial;
		}
		h1 {
			text-align:center;
			position:relative;
		}
		.page {
			width: 80%;
			max-width: 500px;
			margin: auto;
			margin-top: 40px;
			margin-bottom: 40px;
			padding: 80px;
			background: #FFFCF6;
			border: 1px solid #eee;
			box-shadow: 0 10px 6px -6px #777;
		}
		ul {
			border-top: 1px dotted #ccc;
			padding: 10px 0 0 0;
			-webkit-padding-start: 0;
		}
		li {
			padding-top: 10px;
		}
		a,
		a:visited {
			text-decoration: none;
			color: #6295b5;
		}
		a:active,
		a:hover {
			color: #7792a2;
		}
		p {
			text-indent: 1em;
		}
	</style>
</html>`
