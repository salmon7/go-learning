package main

import (
	"reflect"
	"fmt"
)

func main() {
	u := User3{"张三", 20}
	v := reflect.ValueOf(u)

	mPrint := v.MethodByName("Print")
	fmt.Println(mPrint.IsValid())

	args := []reflect.Value{reflect.ValueOf("前缀")}
	fmt.Println(mPrint.Call(args))

}

type User3 struct {
	Name string
	Age  int
}

func (u User3) Print(prfix string) {
	fmt.Printf("%s:Name is %s,Age is %d", prfix, u.Name, u.Age)
}

/*
https://www.flysnow.org/2017/06/13/go-in-action-go-reflect.html

MethodByName方法可以让我们根据一个方法名获取一个方法对象，然后我们构建好该方法需要的参数，最后调用Call就达到了动态调用方法的目的。

获取到的方法我们可以使用IsValid 来判断是否可用（存在）。

这里的参数是一个Value类型的数组，所以需要的参数，我们必须要通过ValueOf函数进行转换。
 */