package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {

	str := `{"age":0}`
	var extra struct {
		Name string `json:"name"`
		Age  int    `json:"age,omitempty"`
	}
	err := json.Unmarshal([]byte(str), &extra)
	if err != nil {
		fmt.Printf("%v",err)
		os.Exit(-1)
	}
	fmt.Printf("Unmarshal: %+v\n", extra)

	eb, err :=json.Marshal(extra)
	if err != nil {
		fmt.Printf("%v",err)
		os.Exit(-1)
	}
	fmt.Printf("Marshal: %s\n", eb)

}

/*
omitempty表示当该字段为对应的零值时，序列化的时候不会序列化该字段

需要注意的是，omitempty生效是在json序列化的过程

 */
