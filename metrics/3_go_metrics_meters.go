package main

import (
	"github.com/rcrowley/go-metrics"
	"time"
	"log"
	"os"
)

func main() {

	m := metrics.NewMeter()
	metrics.Register("quux", m)
	//m.Mark(1)

	go metrics.Log(metrics.DefaultRegistry,
		1 * time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))


	for true {
		time.Sleep(time.Second * 1)
		m.Mark(1)
	}
}

/*
from https://zhuanlan.zhihu.com/p/30441529
 */
