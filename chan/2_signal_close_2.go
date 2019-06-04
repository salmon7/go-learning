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
				signal.Stop(sig)
			} else {
				fmt.Printf("sig is closed\n")
			}
		}
	}
}

/* 可以发送多次 SIGINT 给程序
kill -s SIGINT pid

第一次接收到SIGINT会走代码设定case逻辑，因为有signal.Stop()，所以第二次接收到SIGINT时会走进程的默认逻辑，即退出程序。

所以在监听操作系统的SIGINT或类似的信号下，尽量不要使用signal.Stop()
 */
