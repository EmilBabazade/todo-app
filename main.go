package main

import (
	"fmt"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Wekcome to todo app"))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)

	url := "127.0.0.1:5000"
	fmt.Printf("Serving at %v", url)

	err := http.ListenAndServe(url, mux)
	if err != nil {
		log.Fatal(err.Error())
	}
}
