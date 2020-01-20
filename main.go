package main

import (
				"fmt"
				"log"
				"net/http"
				"github.com/gorilla/mux"
)

func homeLink(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homeLink)
	log.Fatal(http.ListenAndServe(":8081", router))
	// fmt.Println("Hello World!")
}