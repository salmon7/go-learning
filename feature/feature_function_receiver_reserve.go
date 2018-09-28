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
	//router := web.New(Context{}).Middleware((*Context).CSRFMiddleware)
	xx :=(*Context).CSRFMiddleware
	c := Context{}
	xx(&c)

}
