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
*/
