package main

import "fmt"

const (
	a = iota
	b
	c
)

type Updater func() map[string]bool


func MyMethod() Updater{
	return func() map[string]bool{
		ips := make(map[string]bool)
		ips["a"] = true
		return ips
	}
}


func main(){

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(MyMethod()())


}