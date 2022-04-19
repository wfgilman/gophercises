package main

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"example.com/handler"
	// "github.com/boltdb/bolt"
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

	// db, err := bolt.Open("svr.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	// if err != nil {
	// 	fmt.Errorf("could not open db: %s", err)
	// }
	// defer db.Close()
	//
	// db.Update(func(tx *bolt.Tx) error {
	// 	_, err := tx.CreateBucket([]byte("PathUrl"))
	// 	if err != nil {
	// 		return fmt.Errorf("create bucket: %s", err)
	// 	}
	// 	return nil
	// })
	//
	// db.Update(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("PathUrl"))
	// 	err := b.Put([]byte("/disney"), []byte("https://www.disney.com"))
	// 	return err
	// })

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
	// dbHandler, err := handler.DbHandler(db, mux)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", *port),
		Handler: jsonHandler,
	}

	go start(server)

	stopCh, closeChanF := createChannel()
	defer closeChanF()
	log.Println("notified:", <-stopCh)

	shutown(context.Background(), server)
}

func start(server *http.Server) {
	log.Println("server started...")
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	} else {
		log.Println("server stopped gracefully")
	}
}

func shutown(ctx context.Context, server *http.Server) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	} else {
		log.Println("server shutdown")
	}
}

func createChannel() (chan os.Signal, func()) {
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	return stopChan, func() {
		close(stopChan)
	}
}
