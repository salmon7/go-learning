package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := 2

	// 注意这里必须传地址
	v := reflect.ValueOf(&x)
	// 如果上一步传的是地址，这里则是ptr类型
	fmt.Println(v.Kind())
	// 如果上一步传的是地址，这里则为true
	fmt.Println(v.Elem().CanSet())

	v.Elem().SetInt(100)
	fmt.Println(x)
}

/*
from https://www.flysnow.org/2017/06/13/go-in-action-go-reflect.html

以上有几个重点，才可以保证值可以被修改，Value为我们提供了CanSet方法可以帮助我们判断是否可以修改该对象。
 */
