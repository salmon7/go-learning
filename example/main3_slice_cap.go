package main

import "fmt"

func main(){
	x := []int{2, 3, 5, 7, 11}
	y := x[1:3]
	z := y[0:4]
	fmt.Println(y)
	fmt.Println(z)
	/*
	fmt.Println(cap(x))
    fmt.Println(cap(y))
    fmt.Println(cap(z))
	*/
}

/*
Output:
[3 5]
[3 5 7 11]

解释:
可以看到y的长度只有2，z的长度却为4,，一般情况下会认为越界。但是go实现的slice的容量出现使这种情况成为了可能。

分清长度和容量
长度是下标操作的上界，如x[i]中i必须小于长度。而容量是分割操作的上界，如x[i:j]中j不能大于容量。

cap的计算方式
cap = 父cap-子firstindex

*/