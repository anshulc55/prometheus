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

var REQUEST_RESPOND_TIME = promauto.NewSummaryVec(prometheus.SummaryOpts{
	Name: "go_app_response_latency_seconds",
	Help: "Response latency in seconds.",
}, []string{"path"})

func main() {
	// Start the application
	startMyApp()
}

func routeMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start_time := time.Now()
		route := mux.CurrentRoute(r)
		path, _ := route.GetPathTemplate()

		next.ServeHTTP(w, r)
		time_taken := time.Since(start_time)
		REQUEST_RESPOND_TIME.WithLabelValues(path).Observe(time_taken.Seconds())
	})

}

func startMyApp() {
	router := mux.NewRouter()
	router.HandleFunc("/user/{name}", func(rw http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		name := vars["name"]
		greetings := fmt.Sprintf("Hi Buddy, your name is %s :)", name)
		time.Sleep(3 * time.Second)
		rw.Write([]byte(greetings))
	}).Methods("GET")

	router.Use(routeMiddleware)
	log.Println("Starting the application server, Summary Metrics...")
	router.Path("/metrics").Handler(promhttp.Handler())
	http.ListenAndServe("0.0.0.0:8000", router)
}
