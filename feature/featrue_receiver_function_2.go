package main

import "fmt"

type A interface {
	method()
}

type B struct {
	x *int
}

func (b *B) method() {
	b.x = new(int)
	*b.x = 2
}

func change(a A) {
	a.method()
}

func main() {
	b := B{}
	change(&b)
	fmt.Printf("%v", *b.x)
}
/*
Output:
2

解释:
结构体B实现接口A的method方法时,接收者为 变量指针 类型,因此传递给 change 方法中传递该接收者作为接口类型参数时也必须为指针

* 如果实现接口的方法的接收器是指针类型，则传参给接口类型时必需是指针，如果不是，则随便传

from: https://blog.csdn.net/zengming00/article/details/78971147
*/