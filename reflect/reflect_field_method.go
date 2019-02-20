package main

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
)

type User2 struct {
	Name string `json:"name" schema:"schema_name"`
	Age  int    `json:"age" schema:"schema_age"`
}

func (User2) MyMethod_1() {}

func (User2) myMethod_2() {}

func main() {
	var u User2
	h := `{"name":"张三","age":15}`
	err := json.Unmarshal([]byte(h), &u)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	fmt.Printf("%+v\n", u)

	fmt.Println("======================")
	t := reflect.TypeOf(u)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		fmt.Println(sf.Tag)
	}

	fmt.Println("======================")
	t2 := reflect.TypeOf(u)
	for i := 0; i < t2.NumField(); i++ {
		sf := t2.Field(i)
		fmt.Println(sf.Name, sf.Type, sf.Tag.Get("schema"))
	}

	fmt.Println("======================")
	t3 := reflect.TypeOf(u)
	for i := 0; i < t3.NumMethod(); i++ {
		sf := t3.Method(i)
		fmt.Println(sf.Name, sf.Type)
	}
}

/*
from https://www.flysnow.org/2017/06/25/go-in-action-struct-tag.html

多个Key使用空格进行分开，然后使用Get方法获取不同Key的值。
 */
