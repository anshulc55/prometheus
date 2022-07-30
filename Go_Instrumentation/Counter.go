package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var REQUEST_COUNT = promauto.NewCounter(prometheus.CounterOpts{
	Name: "go_app_requests_count",
	Help: "Total App HTTP Requests count.",
})

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
		REQUEST_COUNT.Inc()
	}).Methods("GET")

	log.Println("Starting the application server, Counter Metrics...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8000", router)
}
