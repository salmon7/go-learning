package main

import (
	"encoding/json"
	"fmt"
)

func main(){
	var testJsonNumber struct {
		Number json.Number `json:"number"`
	}
	testJsonNumber.Number = "123.123"
	x := testJsonNumber.Number.String()
	fmt.Println(x)

	jbyte, err := json.Marshal(testJsonNumber)
	if err !=nil{
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n",jbyte)

	var testString struct {
		Number string `json:"number"`
	}
	err = json.Unmarshal(jbyte,&testString)
	if err !=nil{
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%v\n",testString)

}
/*
json.Number的使用示例
- marshal时，可以从string转为number
- unmarshal时，string不能转为number
- 转为json后，其实就是一个number
 */