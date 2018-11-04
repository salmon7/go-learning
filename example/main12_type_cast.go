package main

import "fmt"

func main() {
	var i interface{} = 123

	val,ok := i.(int)
	if ok{
		fmt.Printf("%T->%d\n", val, val)
	}else{
		fmt.Println("not match")
	}
}

/*
类型转换
接口转具体类型：[interface].(int)
*/
