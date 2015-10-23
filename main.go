package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/bugsnag/bugsnag-go"
)

var (
	port int
)

func init() {
	bugsnag.Configure(bugsnag.Configuration{
		APIKey:       os.Getenv("BUGSNAG_API_KEY"),
		ReleaseStage: os.Getenv("BUGSNAG_RELEASE_STAGE"),
	})

	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)

	http.HandleFunc("/locations.json", LocationsHandler)
	http.HandleFunc("/", StaticHandler)

	handler := NewLoggingMiddleware(bugsnag.Handler(nil))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", port), handler); err != nil {
		panic("Error starting!")
	}
}
