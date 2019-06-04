package main

import (
	"fmt"
)

func main() {

	c := make(chan int, 1)
	// c <- 1
	close(c)

	select {
	case x, ok := <-c:
		// 非阻塞。
		// 1.传值给c；
		// 2.关闭信道c
		if ok {
			fmt.Printf("receive c: %v", x)
		} else {
			// x will be zero
			fmt.Printf("c has been closed")
		}
	}

}

/* 可以看到：
1.如果channel c已经被关闭，不但可以读取出已发送的数据，还可以不断的读取零值
2.如果channel c已经被关闭，通过i, ok := <-c可以查看channel的状态，通过判断ok，可知i值是正常读取的值，还是因为关闭返回的零值
2.如果channel c已经被关闭，select会立即返回。

ps:
1.如果channel c已经被关闭，通过range读取，channel关闭后for循环会跳出
2.如果channel c已经被关闭，继续往它发送数据会导致panic: send on closed channel
3.注意switch和select的使用

*/
