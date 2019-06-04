package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	// c := make(chan int, 1)
	defer close(c)
	go func() {
		fmt.Println("send 7")
		c <- 3 + 4
		fmt.Println("after send 7")

		time.Sleep(1 * time.Second)

		fmt.Println("send 8")
		c <- 3 + 5
		fmt.Println("after send 8")
	}()
	i := <-c
	fmt.Println(i)
	time.Sleep(10 * time.Second)
}
/*
1.chan阻塞模式： c := make(chan int)
程序输出：send 7
after send 7
7
send 8

2.chan非阻塞模式（注释第9行，不注释第10行）： c := make(chan int, 1)
程序输出：
send 7
after send 7
7
send 8
after send 8

 */