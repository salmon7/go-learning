package main

import (
	"fmt"
	"reflect"
)

type param struct {
	Ctx *Context
}

// Context ctx struct
type Context struct {
	name string
}

func (c Context) String() string {
	return fmt.Sprintf("my_context")
}

func main() {
	p := param{}
	p.Ctx = &Context{name: "123"}

	fmt.Printf(fmt.Sprintf("%+v\n", p))

	fmt.Println("=====first layer of printValue() depth=0, struct =====")

	r := reflect.ValueOf(p)
	k := r.Kind()
	fmt.Println(k)
	fmt.Println(r.NumField())
	fmt.Println(r.Type().Field(0).Name)

	fmt.Println("=====second layer of printValue() depth=1, ptr =====")

	r0 := r.Field(0)
	k0 := r0.Kind()
	fmt.Println(k0)
	fmt.Println(r0.IsValid())
	fmt.Println(r0.CanInterface())
	fmt.Println(r0.Pointer())
	fmt.Println(r0.Elem().Kind())
	if r0.Elem().IsValid() && r0.Elem().CanInterface() {
		fmt.Println("=====transfer prt to struct =====")
		fmt.Println(r0.Elem().Interface())
		fmt.Println(reflect.ValueOf(r0.Elem().Interface()).Kind())
		var i interface{}
		i = reflect.ValueOf(r0.Elem().Interface())
		switch v := i.(type) {
		case error:
			fmt.Println("ctx is err")
		case fmt.Stringer:
			fmt.Println("ctx is stringer")
			fmt.Println(v)
		}
	} else {
		fmt.Println("can not interface()")
	}
}

/*

前言：

最近线上遇到server输出context的情况，情况与以下链接一致。代码也可在本包server_context_bug.go中查看。
http://xiaorui.cc/2019/06/26/golang-net-http%e8%be%93%e5%87%bacontext%e5%bc%95%e8%b5%b7%e7%9a%84map-panic/

让我觉得奇怪的是，只要把params.Ctx 改成 params.ctx 就不会panic，并且只会输出地址，故才有此文对fmt的相关机制进行探究。


正文：

改变param中Ctx成员的首字母大小写，观察第25行输出情况:
当为Ctx时，输出{Ctx:my_context}
当为ctx时，输出{ctx:0xc42007a1c0}

1.跟踪Ctx时，fmt包调用的情况，核心方法为printArg()，printValue()和handleMethod()

调用printArg()时，将断言至default类别，handlerMethoc()返回false，第一次调用printValue(), depth=0。
在printValue()中将断言至Struct，获取第一个成员变量（即Ctx），并且再次调用printValue(),depth=1。
value的kind()为ptr，depth > 0 && value.IsValid() && value.CanInterface()返回true, 并且Interface()方法返回一个struct（即Ctx）。
再调用handleMethod()方法，由于Ctx实现了String()方法，故在该方法中输出my_context，并且返回true


2.跟踪ctx时，fmt包调用的情况，核心方法为printArg()和printValue()

调用printArg()时，将断言至default类别，handlerMethoc()返回false，第一次调用printValue(), depth=0。
在printValue()中将断言至Struct，获取第一个成员变量（即ctx），并且再次调用printValue(),depth=1。
value的kind()为ptr，depth > 0 && value.IsValid() && value.CanInterface()返回false, 断言至Ptr，调用fmtPointer输出指针。


可以看到关键区别在于大小写是否可见的问题：
如果不可导出（即成员变量首字母小写），则无法把ptr映射成struct，最终只能输出地址；
如果可导出（即成员变量首字母大写），则可以把ptr映射成struct，最终会根据String()方法输出；


3.可以试试把Context实现的String()方法注释掉，会输出什么样的数据？

4.总结
ctx，非指针，实现String()方法   {ctx:{name:123}}
ctx，指针，实现String()方法     {ctx:0xc42007a1c0}
Ctx，非指针，实现String()方法  {Ctx:my_context}
Ctx，指针，实现String()方法     {Ctx:my_context}

ctx，非指针，不实现String()方法   {ctx:{name:123}}
ctx，指针，不实现String()方法      {ctx:0xc42007a1c0}
Ctx，非指针，不实现String()方法  {Ctx:{name:123}}
Ctx，指针，不实现String()方法     {Ctx:0xc42000e1e0}

非指针: struct --> struct --> 首字母大写，value.IsValid() && value.CanInterface() 返回true --> 按是否实现String()输出
               --> struct --> 首字母小写，value.IsValid() && value.CanInterface() 返回false --> 按照struct变量输出

指针: struct --> ptr --> 首字母大写，value.IsValid() && value.CanInterface() 返回true --> 按是否实现String()输出
             --> ptr --> 首字母小写，value.IsValid() && value.CanInterface() 返回false，即首字母小写 --> 按照ptr变量输出

fmt包中也有一句话（https://golang.org/pkg/fmt/）：
When printing a struct, fmt cannot and therefore does not invoke formatting methods such as Error or String on unexported fields.

*/
