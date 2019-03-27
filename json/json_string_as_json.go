package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type TestJson struct {
		Number int    `json:"number"`
		Data   string `json:"data"`
	}
	var json1 TestJson
	json1.Number = 1
	json1.Data = "{\"age\":10}"

	jbyte, err := json.Marshal(json1)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%s\n", jbyte)

	var json2 struct {
		Number int `json:"number"`
		Data struct {
			Age int `json:"age"`
		} `json:"data"`
	}
	err = json.Unmarshal(jbyte, &json2)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	fmt.Printf("%+v\n", json2)

}

/*
验证能否把某个string的json转为json对象
 */
