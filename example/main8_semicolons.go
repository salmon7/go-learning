package main

import "fmt"

func main(){
	a := [10]int{1}

	for _,j := range a {
		fmt.Println(j)
	}

}