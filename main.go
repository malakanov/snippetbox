package main

import (
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)

	log.Println("Server start in on :4000")

	err := http.ListenAndServe("localhost:4000", mux)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hello world"))
	if err != nil {
		log.Fatal(err)
	}
}
