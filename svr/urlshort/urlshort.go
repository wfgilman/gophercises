package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"example.com/handler"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome Home!\n")
}

func main() {
	port := flag.String("port", "8080", "port to start server on")
	flag.Parse()

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/hello", hello)
	mux.HandleFunc("/headers", headers)

	// pathsToUrls := map[string]string{
	// 	"/google": "https://google.com",
	// }

	// yaml := `
	// - path: /google
	//   url: https://google.com
	// - path: /cnn
	//   url: https://www.cnn.com/
	// `

	json := `[
      {
        "path":"/erlang",
        "url":"https://www.erlang.org/"
      },
      {
        "path":"/hacker",
        "url":"https://news.ycombinator.com/"
      }
    ]
  `

	// mapHandler := handler.MapHandler(pathsToUrls, mux)
	// yamlHandler, err := handler.YamlHandler([]byte(yaml), mux)
	jsonHandler, err := handler.JsonHandler([]byte(json), mux)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), jsonHandler))
}
