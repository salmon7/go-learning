package main

import (
	"time"
	"os/signal"
	"os"
	"syscall"
	"fmt"
)

func main() {
	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	for {
		select {
		case rec, ok := <-sig:
			if ok {
				fmt.Printf("receive sig: %v \n", rec)
				time.Sleep(5 * time.Second)
				fmt.Printf("after sleep\n")
			} else {
				fmt.Printf("sig is closed\n")
			}
		}
	}
}

/* 可以发送多次 SIGINT 给程序
kill -s SIGINT pid
 */
