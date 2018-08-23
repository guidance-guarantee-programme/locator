package main

import (
	"fmt"
	"io/ioutil"
	"mime"
	"net/http"
	"html/template"
	"os"
	"path"
	"strings"
)

type EnvConfig struct {
	GoogleMapsApiKey string
}

var getEnv = os.Getenv

func StaticHandler(w http.ResponseWriter, r *http.Request) {
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
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// set MIME type from file extension if possible
	split := strings.Split(filename, ".")
	extension := fmt.Sprintf(".%s", split[len(split)-1])
	if mime := mime.TypeByExtension(extension); mime != "" {
		w.Header().Set("Content-Type", mime)
	}
	// send the contents of the file
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(fmt.Sprintf("Can't read %s", filename))
	}

	t := template.New(filename)
	t, _ = t.Parse(string(contents[:]))
	envConfig := EnvConfig{GoogleMapsApiKey: getEnv("GOOGLE_MAP_API_KEY")}
	t.Execute(w, envConfig)
}
