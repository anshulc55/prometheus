package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Hi Buddy, your name is %s :)", name)
		rw.Write([]byte(greetings))
	}).Methods("GET")

	log.Println("Starting the application server...")
	http.ListenAndServe("0.0.0.0:8000", router)
}
