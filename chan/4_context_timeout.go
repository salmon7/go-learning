package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	context, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	sig := make(chan os.Signal)
	signal.Notify(sig, syscall.SIGINT)

	select {
	case <-sig:
		fmt.Printf("receive sig %v \n", sig)
		// signal.Stop(sig)
		time.Sleep(5 * time.Second)
		fmt.Printf("after sleep\n")
	case <-context.Done():
		fmt.Printf("receive context.Done()\n", )
		cancel()
		fmt.Printf("cancel\n")
	}

}

/*
如果把地21行的注释去掉，程序启动后发送多个SIGINT信号会怎样（kill -s SIGINT pid）？参考2_signal_close_2.go


 */
