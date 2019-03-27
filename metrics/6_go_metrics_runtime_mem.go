package main

import (
	"github.com/rcrowley/go-metrics"
	"time"
	"log"
	"os"
	"sync"
)

func main() {
	metrics.RegisterRuntimeMemStats(metrics.DefaultRegistry)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	go metrics.CaptureRuntimeMemStats(metrics.DefaultRegistry, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
