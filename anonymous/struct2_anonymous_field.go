package main

type A1 struct {
	ax int
	ay int
	B1
}
type B1 struct {
	bx float32
	by float32
}

type A2 struct {
	ax int
	ay int
	b  B2
}
type B2 struct {
	bx float32
	by float32
}

type A3 struct {
	B3
	C3
}
type B3 struct {
	x float32
	y float32
}
type C3 struct {
	x float32
	y float32
}

func main() {
	// 如果A包含B，那么如果以A1、B1的关系，则称B1为匿名结构体。使用或者初始化bx/by变量时，可以通过 a.B.bx/by 或者 a.bx/by 使用
	a1 := A1{1,2,B1{3,4}}
	a1.ax = 1
	a1.ay = 2
	a1.bx = 3
	a1.by = 4
	a1.B1.bx = 3
	a1.B1.by = 4

	// 如果A包含B，那么如果以A2、B2的关系，则B2不是匿名结构体。使用或者初始化bx/by变量时，必须通过b使用，如 a.b.bx/by。需要注意的是，如果使用这种方式，则A不能直接使用B的方法
	a2 := A2{1,2,B2{3,4}}
	a2.ax = 1
	a2.ay = 2
	// a2.bx = 3
	// a2.by = 4
	a2.b.bx = 3
	a2.b.by = 4

	// 如果A包含B和C，那么如果以A3、B3、C3的关系，使用或者初始化a.x/a.y变量时，则会编译报错，如果不使用的话，编译器不会发现这个问题
	// a3 := A3{B3{1,2},C3{3,4}}
	// a3.x = 2 // 编译报错: anonymous/struct2_anonymous_field.go:57:4: ambiguous selector a3.x
}

/*
参考：https://wiki.jikexueyuan.com/project/the-way-to-go/10.5.html
 */