package main

import "fmt"

func main() {
	testSwitch('1')
}

func testSwitch(c byte) {
	switch {
	case '0' <= c && c <= '9':
		fmt.Println("This is a number")
	case 'a' <= c && c <= 'z':
		fmt.Println("This is a alp")
	case 'A' <= c && c <= 'Z':
		fmt.Println("This is a upper alp")
	default:
		fmt.Println("This is not accepted")
	}
}
/*

与c不同的是，go的switch中的case子句并不会因为没有break就一直执行，而是只执行一个case
并且switch中的break是为了提前结束case后的代码，从而跳转到switch语句块后


*/