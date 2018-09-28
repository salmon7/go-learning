package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	x, y := 0, 1
	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 5; i++ {
		fmt.Println(f())
	}
}
/*
from https://blog.csdn.net/wangshubo1989/article/details/79217291
 */