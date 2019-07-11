package main

import (
	"context"
	"fmt"
	"net/http"
)

type params struct {
	Ctx context.Context
}

func panic(w http.ResponseWriter, r *http.Request) {
	p := params{}
	p.Ctx = r.Context()
	fmt.Printf("%+v\n", p)

	/*
	switch f := p.ctx.(type) {
	default:
		val := reflect.ValueOf(f)
		fmt.Printf("val=%v\n", val)
		kind := val.Kind()
		fmt.Printf("val.kind=%v\n", kind)
		fmt.Printf("val.Pointer=%v\n", val.Pointer())
		a := val.Elem()
		fmt.Printf("val.Elem.Kind=%v\n", a.Kind())
	}*/

}

func main() {
	http.HandleFunc("/", panic)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Println(err)
	}
}