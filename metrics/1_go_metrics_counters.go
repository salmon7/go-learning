package main

import (
	"time"
	"os"
	"log"
	"github.com/rcrowley/go-metrics"
)

func main() {
	c := metrics.NewCounter()
	metrics.Register("foo", c)

	go metrics.Log(metrics.DefaultRegistry,
		1*time.Second,
		log.New(os.Stdout, "metrics: ", log.Lmicroseconds))

	var j int64
	j = 2
	for true {
		time.Sleep(time.Second * 2)
		c.Inc(2)
		j++
	}
}

/*
from https://zhuanlan.zhihu.com/p/30441529
 */
