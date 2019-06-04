package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	context, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// cancel()

	select {
	case <-context.Done():
		// context.Done()返回的是closed的chan
		fmt.Printf("context done: %v", context.Err())
	}
}

/*

等待两秒后，输出：
context done: context deadline exceeded

12行不注释，立即输出：
context done: context canceled

*/
