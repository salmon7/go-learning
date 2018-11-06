package main

import (
	"fmt"
)

type tag struct {
	value int32
}

func (_tag tag) Change() {
	_tag.value = int32(987)
}

type tag2 struct {
	value int32
}

func (_tag *tag2) Change2() {
	_tag.value = int32(987)
}

func main() {
	_tag := new(tag)
	_tag.value = 123
	_tag.Change()
	fmt.Println(_tag)
	_tag.Change()
	fmt.Println(_tag)

	_tag2 := tag2{41}
	_tag2.Change2()
	fmt.Println(_tag2)
	_tag2.Change2()
	fmt.Println(_tag2)

}
/*
Output:
&{123}
&{123}
{987}
{987}

解释:
在main中_tag为 *tag 类型变量，_tag2为 tag2类型变量，它们对应Change和Change2方法的接受者为值和指针。

* 如果想要修改struct的成员的值，method被定义时候其ReceiverType必须是struct*形式。如果ReceiverType是struct，则无法改变struct成员的值。
* 记住，这与结构体变量本身是 struct 类型，还是 *struct 类型没有关系

from: https://blog.csdn.net/menggucaoyuan/article/details/43056261
*/