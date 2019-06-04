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
				close(sig)
			} else {
				fmt.Printf("sig is closed\n")
			}
		}
	}
}

/* 可以发送多次 SIGINT 给程序
kill -s SIGINT pid

第一次接收到SIGINT会走代码设定case逻辑，在此之后由于close()掉了sig，所以ok==false一直输出“sig is closed”，并且再次发送SIGINT时程序会崩溃：
panic: send on closed channel

所以在监听操作系统的信号下，尽量不要使用close()
 */
