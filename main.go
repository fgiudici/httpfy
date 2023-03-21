package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var port int
	var dir string

	flag.IntVar(&port, "p", 8080, "listening port")
	flag.StringVar(&dir, "d", "", "serving dir"+" (default current dir)")

	flag.Parse()
	if dir == "" {
		var err error
		dir, err = os.Getwd()
		if err != nil {
			log.Fatal(err)
		}
	}

	fileinfo, err := os.Stat(dir)
	if err != nil {
		log.Fatal(err)
	}
	if !fileinfo.IsDir() {
		log.Fatalf("%s is not a directory", dir)
	}

	log.Printf("Serving '%s' on port %d", dir, port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), http.FileServer(http.Dir(dir))))
}
