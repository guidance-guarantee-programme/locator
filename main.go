package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)

	http.HandleFunc("/locations.json", LocationsHandler)
	http.HandleFunc("/", StaticHandler)

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}