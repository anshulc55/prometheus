package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var REQUEST_INPROGRESS = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "go_app_requests_inprogress",
	Help: "Number of application requests in progress",
})

func main() {
	// Start the application
	startMyApp()
}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", func(rw http.ResponseWriter, r *http.Request) {
		REQUEST_INPROGRESS.Inc()
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Hi Buddy, your name is %s :)", name)
		time.Sleep(5 * time.Second)
		rw.Write([]byte(greetings))

		REQUEST_INPROGRESS.Dec()
	}).Methods("GET")

	log.Println("Starting the application server, Gauge Metrics...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8000", router)
}
