package main

import "fmt"

func main()	{
	a := new([10]int)
	fmt.Println(a)

	b := new([]int)
	//*b = append(*b,1)
	fmt.Println(b)

	c := make([]int,10)
	fmt.Println(c)
}

/*
a是数组，b、c是切片，如果不能理解请参考 mian4_array.go

Output:
&[0 0 0 0 0 0 0 0 0 0]
&[]
[0 0 0 0 0 0 0 0 0 0]

解释:
Go有两个数据结构创建函数：new和make。两者的区别在学习Go语言的初期是一个常见的混淆点。
基本的区别是new(T)返回一个*T，而make(T, args)返回一个普通的T。
通常情况下，T内部有一些隐式的指针（图中的灰色箭头）。一句话，new返回一个指向已清零内存的指针，而make返回一个复杂的结构。
即new返回一个指定类型的指针，而make返回指向指定结构类型的slice数据结构

对于b的赋值：
	b := new([]int)
	*b = append(*b,1)
or
	b := *new([]int)
	b = append(b,1)


*/