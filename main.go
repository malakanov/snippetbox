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

func create(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		writer.WriteHeader(405)
		_, err := writer.Write([]byte("Method are not allowed"))
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	_, err := writer.Write([]byte("Create snippet"))
	if err != nil {
		log.Fatal(err)
	}
}

func detail(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Show a snippet"))
	if err != nil {
		log.Fatal(err)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	_, err := writer.Write([]byte("Hello world"))

	if err != nil {
		log.Fatal(err)
	}
}
