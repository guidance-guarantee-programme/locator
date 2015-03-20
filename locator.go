package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path"
	"strings"
)

var (
	port int
)

func handler(w http.ResponseWriter, r *http.Request) {
	var responseCode int

	// check for requested path under public/
	pwd, _ := os.Getwd()
	filename := path.Join(pwd, "public", r.URL.Path)
	stat, err := os.Stat(filename)

	// append index.html if path is a directory
	if !os.IsNotExist(err) && stat.Mode().IsDir() {
		filename = path.Join(filename, "index.html")
		stat, err = os.Stat(filename)
	}

	if os.IsNotExist(err) {
		responseCode = http.StatusNotFound
		w.WriteHeader(responseCode)
	} else {
		// set MIME type from file extension if possible
		split := strings.Split(filename, ".")
		extension := fmt.Sprintf(".%s", split[len(split)-1])
		if mime := mime.TypeByExtension(extension); mime != "" {
			w.Header().Set("Content-Type", mime)
		}
		// send the contents of the file
		responseCode = http.StatusOK
		contents, _ := ioutil.ReadFile(filename)
		w.Write(contents)
	}

	fmt.Printf("%s %s %d\n", r.Method, r.URL.Path, responseCode)
}

func init() {
	flag.IntVar(&port, "p", 8080, "Port to listen on")
	flag.Parse()
}

func main() {
	fmt.Printf("> Starting on http://0.0.0.0:%d\n", port)
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Println("Error starting!")
	}
}
