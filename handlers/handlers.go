package handlers

import (
	"log"
	"net/http"
	"sync/atomic"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	httpRequestsResponseTime prometheus.Summary
)

func init() {
	httpRequestsResponseTime = prometheus.NewSummary(prometheus.SummaryOpts{
		Namespace: "http",
		Name:      "response_time_seconds",
		Help:      "Request response times",
	})

	prometheus.MustRegister(httpRequestsResponseTime)
}

// Router is the mux Router for the Service
func Router(buildTime, commit, release string) http.Handler {
	isReady := &atomic.Value{}
	isReady.Store(false)
	go func() {
		log.Printf("Readyz probe is negative by default...")
		time.Sleep(10 * time.Second)
		isReady.Store(true)
		log.Printf("Readyz probe is positive.")
	}()

	r := mux.NewRouter()
	r.HandleFunc("/home", home(buildTime, commit, release)).Methods("GET")
	r.HandleFunc("/healthz", healthz).Methods("GET")
	r.HandleFunc("/readyz", readyz(isReady)).Methods("GET")
	r.Handle("/metrics", promhttp.Handler())

	withMetrics := middleware(r)
	return withMetrics
}

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		next.ServeHTTP(w, r)

		httpRequestsResponseTime.Observe(float64(time.Since(start).Seconds()))
	})
}
