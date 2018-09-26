package main

import "fmt"

type xx interface {

}

func main() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}

	var x xx
	switch x := x.(type) {
	default:
		fmt.Printf("unexpected type %T\n", x) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", x) // t has type bool
	case int:
		fmt.Printf("integer %d\n", x) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *x) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *x) // t has type *int
	}

}

func functionOfSomeType() *int {
	a := 1
	return &a
}
/*
".(type)"运算只能用interface类型变量调用
*/