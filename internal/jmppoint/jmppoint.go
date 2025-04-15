package jmppoint

import (
    //"fmt"
    //"net/http"
    //"time"

	//"privcrawler/internal/crawler"
    "github.com/prometheus/client_golang/prometheus"
    //"github.com/prometheus/client_golang/prometheus/promhttp"
)

// ---- DATA STRUCTURES ---- //

//Metrics: testing structure for prometheus metrics
var (
	// httpRequests: Counter for total number of HTTP requests.
    httpRequests = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"path"},
    )

	// requestDuration: Histogram for response time of HTTP requests.
    requestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "Histogram of response time for handler",
            Buckets: prometheus.DefBuckets,
        },
        []string{"path"},
    )
)



// ---- Functions ---- //
 
// Function: Initialize 
// Operation: This will initialize the Prometheus metrics and register them with the default registry.
// Return: None
func init() {
    // Register metrics
    prometheus.MustRegister(httpRequests)
    prometheus.MustRegister(requestDuration)
}

// Function 
