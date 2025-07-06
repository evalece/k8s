package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ( 
	queueBacklog = prometheus.NewGauge(prometheus.GaugeOpts{ //a value at a specific moment (like CPU usage or queue size)
		Name: "queue_backlog_total", // meta data, but later recommended protobuf 
		Help: "Number of tasks currently in the backlog queue",
	})
)

func init() { // executes before main()
	prometheus.MustRegister(queueBacklog) //For Optimization- Quick check to avoid having the OS to load main(), if it's set to fail since it will be a larger process to load
}

func main() {
	go func() {
		for {
			backlog := rand.Intn(200) // Simulated value for backlog requests; https://prometheus.io/docs/introduction/overview/ 
			queueBacklog.Set(float64(backlog))
			log.Printf("[INFO] backlog = %d", backlog)
			time.Sleep(5 * time.Second) // repeat every 5 sec
		}
	}()
	// 
	// http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg})) 
	
	// https://pkg.go.dev/github.com/prometheus/client_golang/prometheus 
	http.Handle("/metrics", promhttp.Handler())  // Handler: render in prom's pre-defined format at /metrics endpoint 
	port := 8066 // later make known to prom in prom's manifest 
	//log.Printf("Listening on :%d", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
