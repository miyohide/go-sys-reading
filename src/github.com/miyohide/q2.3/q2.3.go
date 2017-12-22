package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")

	source := map[string]string{
		"Hello": "World",
	}
	file, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	gz := gzip.NewWriter(w)
	defer gz.Close()

	writer := io.MultiWriter(file, os.Stdout, gz)

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
