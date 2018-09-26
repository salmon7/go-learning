package main

import (
	"fmt"
)

type Handler interface {
	ServerHTTP(string)
}

type HandlerFunc func(string)

func (f HandlerFunc) ServeHTTP(s string) {
	f(s)
}

type OldCaller struct{

}

func (OldCaller) OtherFunc(other string) HandlerFunc {
	return func(s string){
		fmt.Println("This is in OtherFunc")
		fmt.Println(s)
		fmt.Println("other:"+other)
	}
}

func (OldCaller) ServerHTTP(s string){
	fmt.Println("This is in ServerHTTP")
	fmt.Println(s)
}

type Caller struct{

}

func (Caller) ServerHTTP(s string){
	fmt.Println("This is in ServerHTTP")
	fmt.Println(s)
}

func CallTheHandlerFunc(handlerFunc HandlerFunc){
	handlerFunc.ServeHTTP("This is unified portal ")
}

func main(){
	var handler Handler = Caller{}
	handler.ServerHTTP("handler")

	oldCaller := OldCaller{}
	var handlerFunc HandlerFunc =oldCaller.OtherFunc("aaaaaaaaa")
	handlerFunc.ServeHTTP("oldCaller")

	//CallTheHandlerFunc(handlerFunc)
}

