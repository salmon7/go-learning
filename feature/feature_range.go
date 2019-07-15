package main

import (
	"fmt"
	"time"
)

func main() {
	b := make([]string, 4)
	b[0] = "1"
	b[1] = "2"
	b[2] = "3"
	b[3] = "4"

	fmt.Printf("%p\n", b)
	go func() {
		// c := b
		for {
			for k, v := range b {
				x := b[k]
				fmt.Printf("go1 %d %s %p %p\n", k, v, &x, b)
				time.Sleep(1 * time.Second)
			}
		}
	}()

	time.Sleep(5 * time.Second)

	a := make([]string, 5)
	a[0] = "a"
	a[1] = "b"
	a[2] = "c"
	a[3] = "d"
	a[4] = "e"

	b = a
	//b = nil
	for k, v := range b {
		x := b[k]
		fmt.Printf("go2 %d %s %p %p\n", k, v, &x, b)
	}

	time.Sleep(10 * time.Second)

}

/*
最近在发现线上fmt.Println(req.context)会panic问题后，对map的操作总是很小心。这两天又接到新需求，需要对map进行操作。

对于range map操作，大家应知道go有对应的语法糖，对于range出来的value是临时变量，如果要更改某个map的key对应的值或者某个slice的index对应的值，应该通过map[key]或者slice[index]去更改，而不是改value本身。

还有一个问题是，如果map/slice本身的地址已经被改变了呢？

如上面代码中在goroutine遍历b时，36行代码改变了b的地址，那么goroutine中for range还是原来的那个么？

先说结论，遍历的还是原来map/slice。用slice尝试一下（用map应该也一样，只是map的遍历顺序每次都不一样，不太好复现）。

程序输出：
0xc42005a040
go1 0 1 0xc42008e000 0xc42005a040
go1 1 2 0xc42000e1e0 0xc42005a040
go1 2 3 0xc4200a0000 0xc42005a040
go1 3 4 0xc42000e210 0xc42005a040
go1 0 1 0xc4200a0020 0xc42005a040   // 注意1
go2 0 a 0xc42000e230 0xc4200820f0
go2 1 b 0xc42000e250 0xc4200820f0
go2 2 c 0xc42000e270 0xc4200820f0
go2 3 d 0xc42000e290 0xc4200820f0
go2 4 e 0xc42000e2b0 0xc4200820f0
go1 1 2 0xc4200a0040 0xc4200820f0   // 注意2
go1 2 3 0xc42000e2d0 0xc4200820f0
go1 3 4 0xc4200a0060 0xc4200820f0
go1 0 a 0xc42000e2f0 0xc4200820f0
go1 1 b 0xc4200a0080 0xc4200820f0
go1 2 c 0xc42000e310 0xc4200820f0
go1 3 d 0xc4200a00a0 0xc4200820f0
go1 4 e 0xc42000e330 0xc4200820f0
go1 0 a 0xc4200a00c0 0xc4200820f0
go1 1 b 0xc42000e350 0xc4200820f0

可以看到在【注意1】和【注意2】之间b的地址已经被改变，但是在goroutine中的一个range循环时读取的依然是一个旧的slice，这次range跨越了b的指针变化，但是goroutine输出的值依然是旧的slice。需要注意的是最后一个列输出变为0xc4200820f0病不代表range的就是最新的,可以把第17行的注释去掉，并且groutine遍历时改为遍历c，程序输出:

0xc420058040
go1 0 1 0xc42008c000 0xc420058040
go1 1 2 0xc42009e000 0xc420058040
go1 2 3 0xc42000e1e0 0xc420058040
go1 3 4 0xc42009e020 0xc420058040
go1 0 1 0xc42000e210 0xc420058040
go2 0 a 0xc42009e060 0xc4200ae000
go2 1 b 0xc42009e080 0xc4200ae000
go2 2 c 0xc42009e0a0 0xc4200ae000
go2 3 d 0xc42009e0c0 0xc4200ae000
go2 4 e 0xc42009e0e0 0xc4200ae000
go1 1 2 0xc42009e040 0xc420058040
go1 2 3 0xc42000e230 0xc420058040
go1 3 4 0xc42008c030 0xc420058040
go1 0 1 0xc42000e250 0xc420058040
go1 1 2 0xc42008c050 0xc420058040
go1 2 3 0xc42000e270 0xc420058040
go1 3 4 0xc42008c070 0xc420058040
go1 0 1 0xc42000e290 0xc420058040
go1 1 2 0xc42008c090 0xc420058040
go1 2 3 0xc42000e2b0 0xc420058040

可以看到只要把b的地址先存到另一个变量，for range 的时候只要range这个变量，就不会受到外面的影响。其实，go在for range之前的已经帮我们做了一步变量的赋值，因此c:=b是多此一举。


这里有更详细的内容：
http://www.voidcn.com/article/p-yrxdfibg-boy.html
https://draveness.me/golang/keyword/golang-for-range.html

*/
