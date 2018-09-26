package main

import "fmt"

func main() {
	a := [10]int{1}

	for j, i := 0, 0; i < len(a); i, j = i+1, 1 {
		fmt.Println(a[i])
		fmt.Println(j)
	}
}
