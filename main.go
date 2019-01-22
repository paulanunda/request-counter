package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	requestsServed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "request_counter_requests_served_total",
		Help: "The total number of requests served",
	})
)

func main() {
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/increment", func(w http.ResponseWriter, r *http.Request) {
		requestsServed.Inc()
		log.Info("Request served for '/'")
		fmt.Fprintf(w, "This is a website server by a Go HTTP server.")
	})

	http.ListenAndServe(":3001", nil)
}
