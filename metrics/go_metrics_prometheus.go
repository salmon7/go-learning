package main

import (
	"github.com/deathowl/go-metrics-prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/rcrowley/go-metrics"
	"time"
	"fmt"
	"log"
)

func main() {
	prometheusRegistry := prometheus.NewRegistry()
	metricsRegistry := metrics.NewRegistry()

	pClient := prometheusmetrics.NewPrometheusProvider(metricsRegistry, "test", "subsys", prometheusRegistry, 1*time.Second)
	gm := metrics.NewGauge()
	metricsRegistry.Register("gauge", gm)
	gm.Update(2)
	go pClient.UpdatePrometheusMetrics()
	gm.Update(13)
	time.Sleep(5 * time.Second)
	metrics, _ := prometheusRegistry.Gather()
	if len(metrics) == 0 {
		log.Fatalf("prometheus was unable to register the metric")
	}

	serialized := fmt.Sprint(metrics[0])
	expected := fmt.Sprintf("name:\"test_subsys_gauge\" help:\"gauge\" type:GAUGE metric:<gauge:<value:%d > > ", gm.Value())
	if serialized != expected {
		log.Printf("serialized %s", serialized)
		log.Printf("expected %s", expected)
		log.Fatalf("Go-metrics value and prometheus metrics value do not match")
	}
}

/*
from https://github.com/deathowl/go-metrics-prometheus/blob/master/prometheusmetrics_test.go
 */