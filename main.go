package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/home", helloWorld).Methods("GET")

	http.ListenAndServe(":8000", router)
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("It works!")
}
