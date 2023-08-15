package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func create(writer http.ResponseWriter, request *http.Request) {
	if request.Method != "POST" {
		writer.Header().Set("Allow", "POST")
		http.Error(writer, "Method not allowed", 405)
		return
	}

	_, err := writer.Write([]byte("Create snippet"))
	if err != nil {
		log.Fatal(err)
	}
}

func detail(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(writer, request)
		return
	}
	fmt.Fprintf(writer, "Display a specific snippet with ID %d...", id)
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
