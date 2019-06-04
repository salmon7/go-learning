package main

import (
	"time"
	"os/signal"
	"os"
	"syscall"
	"fmt"
)

func main() {
	// sig := make(chan os.Signal)
	sig := make(chan os.Signal, 1)
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
		default:
			fmt.Println("default")
			time.Sleep(5 * time.Second)
		}
	}
}

/*
1.select:
如果default存在的情况下，并且没有case需要处理，则会选择default去处理。
如果没有default，则select语句会阻塞，直到某个case需要处理。

2.这里的sig需要有cache存储终端传来信号，因为这个传递过程由signal包控制，如果sig未接收，signal不会重新传递，即必须要在default输出5s后马上
输入信号，select才会可能选择第一个case。所以如果没有cache，程序几乎会一直输出 “default”

3.如果是程序其他方式传递chan（而不是sig传递的chan），则其跟case有同等概率获得执行的机会。

 */
