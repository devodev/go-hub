package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var addr string
	flag.StringVar(&addr, "addr", ":8080", "http service address")
	flag.Parse()

	hub := newHub()
	go hub.Run()

	log.Printf("HTTP Server listening on %q", addr)
	err := http.ListenAndServe(addr, hub.Handler())
	if err != nil {
		log.Fatal(err)
	}
}
