package main

import "fmt"

func main(){
	var a []int
	//fmt.Println(a)
	if a == nil {
		fmt.Println("a is nil")
	}else{
		fmt.Println("a not nil")
	}


	b := []int{}
	//fmt.Println(b)
	if b == nil {
		fmt.Println("b is nil")
	}else{
		fmt.Println("b not nil")
	}


	c := new([]int)
	//fmt.Println(c)
	if c == nil {
		fmt.Println("c is nil")
	}else{
		fmt.Println("c not nil")
	}
}


/*
Output:
a is nil
b not nil
c not nil

解释:
a被定义为一个[]int类型的变量（实际为slice），未进行初始化，所以为nil。
b被定义为一个[]int类型的变量（实际为slice），并进行初始化，len和cap都为0，所以不为nil。
c是使用new进行创建的，它会创建一个指向[]int类型的指针（实际为slice类型的指针），并进行初始化，len和cap都为0，所以不为nil。

更正：
	c := new([]int)也有错误，应该比较 *c==nil ，实际上并未初始化，并且打印 "c is nil"
*/