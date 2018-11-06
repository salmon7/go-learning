package main

import "fmt"

func main() {

	i := 1
	defer defer_method1(i)
	i = 2

}

func defer_method1(i int) {
	fmt.Println(i)
}

/*
Output:
1

解释:
defer的执行需要分两种情况
* defer所标记的函数用变量传递，即使在defer后该值被改变,函数所接收的值依然不变
* defer所标记的函数用变量的地址传递，在defer后该值被改变,函数所接收的值跟随着变化

如果上面的方法定义传递换成地址，则输出2
可以这么认为，虽然defer函数最后才被执行，但是用defer标记函数时就已经发生了参数传递，进行了值复制
*/
