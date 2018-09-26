package main

import "fmt"

type Error struct {
	errCode uint8
}

func (e *Error) Error() string {
	switch e.errCode {
	case 1:
		return "file not found"
	case 2:
		return "time out"
	case 3:
		return "permission denied"
	default:
		return "unknown error"
	}
}

func checkError(err error) {
	if err == nil{
		fmt.Println("checkError err is nil")
	}

	if err != nil {
		fmt.Println("checkError err is not nil")
		panic(err)
	}
}
func main() {

	var e *Error
	//e = &Error{errCode:10}

	fmt.Println(e)
	if e==nil{
		fmt.Println("e is nil")
	}

	checkError(e)
}

/*

提示: 注意checkError函数的参数为error，而不是Error或者 *Error

Output:
<nil>
e is nil
checkError err is not nil
fatal error: panic while printing panic value
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x482075]

goroutine 1 [running]:
runtime.throw(0x4b67bc, 0x20) ......

解释:
interface类型在底层的实现是有两个元素的，一个表示存储对象的类型，一个表示存储对象的值

nil在Go语言中，表示零值，而且nil只能赋值给指针、channel、func、interface、map或slice类型的变量。如果未遵循这个规则，则会引发panic。这在Go语言的官方类型介绍中有提及。而对于一个interface类型的对象，只有其内部的类型域和值域均未设置(nil, nil)的情况下才等于nil。https://golang.org/pkg/builtin/#pkg-variables

如何避免：
	避免将一个有可能为nil的具体类型的值赋值给interface变量
	直接返回nil给接口，而不是返回一个自定义的错误
*/