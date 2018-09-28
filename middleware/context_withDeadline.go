package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	defer cancel()

	select {
	case <-time.After(1 * time.Second):
		fmt.Println("overslept")
	case <-ctx.Done():
		if dl,ok :=ctx.Deadline();ok==true{
			fmt.Println(dl)
		}
		fmt.Println(ctx.Err())
	}

}
/*
from https://blog.csdn.net/wangshubo1989/article/details/78910935
 */
