package main

import (
	"github.com/rcrowley/go-metrics"
	"time"
	"log"
	"os"
)

func main() {

	s := metrics.NewExpDecaySample(1024, 0.015) // or metrics.NewUniformSample(1028)

	h := metrics.NewHistogram(s)

	metrics.Register("baz", h)
	//h.Update(1)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 1
	for true {
		time.Sleep(time.Second * 1)
		j++
		h.Update(j)
	}
}

/*
from https://zhuanlan.zhihu.com/p/30441529
 */
