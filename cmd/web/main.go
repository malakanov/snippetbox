package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", detail)
	mux.HandleFunc("/snippet/create", create)

	log.Println("Server start in on :4000")

	err := http.ListenAndServe("localhost:4000", mux)

	if err != nil {
		log.Fatalf(err.Error())
	}
}
