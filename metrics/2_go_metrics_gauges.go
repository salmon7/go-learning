package main

import (
	"time"
	"os"
	"log"
	"github.com/rcrowley/go-metrics"
)

func main() {
	g := metrics.NewGauge()
	metrics.Register("bar", g)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 2)
		g.Update(j)
		j++
	}
}

/*
from https://zhuanlan.zhihu.com/p/30441529
 */
