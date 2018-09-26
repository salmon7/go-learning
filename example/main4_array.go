package main

import "fmt"

func main()	{
	e := new([10]int)
	fmt.Println(e)
}

/*
Output:
&[0 0 0 0 0 0 0 0 0 0]

解释:
	初步结论，凡是指定了具体的长度，则为数组，如var a [10]int，b:=[10]int{}，c:=new([10]int)
	例外：   
	 a := [...]string   { "no error",  "Eio", "invalid argument"}	 	 

*/