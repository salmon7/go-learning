package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	u := User1{"张三", 20}

	// 获取TypeOf和ValueOf
	t := reflect.TypeOf(u)
	fmt.Println(t)
	v := reflect.ValueOf(u)
	fmt.Println(v)

	// 对于以上这两种输出，Go语言还通过fmt.Printf函数为我们提供了简便的方法
	fmt.Println("======================")
	fmt.Printf("%T\n", u)
	fmt.Printf("%v\n", u)

	// 从reflect.Value转原始类型
	fmt.Println("======================")
	u1 := v.Interface().(User1)
	fmt.Println(u1)
	// reflect.Value又同时持有一个对象的reflect.Value和reflect.Type,所以我们可以通过reflect.Value的Interface方法实现还原
	t1 := v.Type()
	fmt.Println(t1)

	// 获取类型底层类型, 可以使用Value对象和Type对象
	fmt.Println("======================")
	fmt.Println(t.Kind())
	fmt.Println(v.Kind())
}

/*
from https://www.flysnow.org/2017/06/13/go-in-action-go-reflect.html

在Go的反射定义中，任何接口都会由两部分组成的，一个是接口的具体类型，一个是具体类型对应的值。
比如var i int = 3 ，因为interface{}可以表示任何类型，所以变量i可以转为一个接口interface{}
这个变量在Go反射中的表示就是<Value,Type>，其中Value为变量的值3,Type变量的为类型int。
 */
