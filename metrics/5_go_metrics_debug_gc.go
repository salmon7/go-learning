package main

import (
	"github.com/rcrowley/go-metrics"
	"time"
	"log"
	"os"
)

func main() {
	metrics.RegisterDebugGCStats(metrics.DefaultRegistry)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	go metrics.CaptureDebugGCStats(metrics.DefaultRegistry, 1)

	ch := make(chan int, 1)
	<-ch
}
