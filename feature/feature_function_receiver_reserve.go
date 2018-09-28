package main

import (
	"fmt"
)

type Context struct {
	CSRFToken string
	User      string
}

func (c *Context) CSRFMiddleware() {
	fmt.Println("xx")
}

func main() {
	xx :=(*Context).CSRFMiddleware
	c := Context{}
	xx(&c)

}
/*
from https://blog.questionable.services/article/map-string-interface/#context-structs
受到例子的启发，能够传递方法
*/