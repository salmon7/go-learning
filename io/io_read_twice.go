package main

import (
	"strings"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"bytes"
)

func main() {

	readAllTwice()
	readAllAndJsonDecode()
	readAllAndJsonDecode2()
}

func readAllTwice() {
	fmt.Println("------readAllTwice------")
	fakeBody := strings.NewReader(`{"name":"name_test","age":20}`)
	bodybytes, err := ioutil.ReadAll(fakeBody)
	if err != nil {
		fmt.Printf("readall err: %v", err)
		return
	}
	fmt.Printf("bodybytes is: %s\n", bodybytes)

	rbody, err := ioutil.ReadAll(fakeBody)
	if err != nil {
		fmt.Printf("again readall err: %v", err)
		return
	}
	fmt.Printf("rbody is: %s\n", rbody)

}

func readAllAndJsonDecode() {
	fmt.Println("------readAllAndJsonDecode------")
	fakeBody := strings.NewReader(`{"name":"name_test","age":20}`)
	bodybytes, err := ioutil.ReadAll(fakeBody)
	if err != nil {
		fmt.Printf("readall err: %v", err)
		return
	}
	fmt.Printf("bodybytes is: %s\n", bodybytes)

	var person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	err = json.NewDecoder(fakeBody).Decode(&person)
	if err != nil {
		fmt.Printf("decode err: %v\n", err)
		return
	}
	fmt.Printf("%#v\n", person)
}

func readAllAndJsonDecode2() {
	fmt.Println("------readAllAndJsonDecode2------")
	fakeBody := strings.NewReader(`{"name":"name_test","age":20}`)

	bodybytes, err := ioutil.ReadAll(fakeBody)

	if err != nil {
		fmt.Printf("readall err: %v", err)
		return
	}
	fmt.Printf("bodybytes is: %s\n", bodybytes)

	var person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	r := ioutil.NopCloser(bytes.NewBuffer(bodybytes))
	err = json.NewDecoder(r).Decode(&person)
	if err != nil {
		fmt.Printf("decode err: %v", err)
		return
	}
	fmt.Printf("%#v\n", person)
}
/*
最近在做一次回调接口的兼容性升级，原本回调方支持 application/json 协议，现在更新为 application/x-www-form-urlencoded 协议，所以现有接口需要兼容以上两种支持

1.r.ParseForm(), 把body解析到 r.Form 或者 r.PostForm 中，再用gorilla.schema包解析到对应的结构体中

2.json.NewDecoder(r.Body).Decode(&bodyjson)，直接读body的二进制流，并解码到对应的结构体中

一般情况下，如果不对body做特殊处理，body只能被读一次，因为body是一个二进制流。开发过程中，先做1后做2，先做2后做1，得到的结果完全不同

先做1后做2：
①如果返回的是 application/json，1不能解析，2可以正常解析
②如果返回的是 application/x-www-form-urlencoded，1能解析，直接返回

先做2后做1：
③如果返回的是 application/json，2可以正常解析，直接返回
④如果返回的是 application/x-www-form-urlencoded，2不能解析，1不能解析 【异常点】

可以看到④的异常点，即使返回的是 application/x-www-form-urlencoded，1居然也不能解析，并且输出r.Form也为空。

给出了几个可能结果：
1.一般body只能被读一次，所以异常情况④是因为body被读取过了，下次再读时返回的body为空，就不能解析到r.Form了。但是，可以看到①也是【读取了两次body】，第一次解析失败，第二次却解析成功了。
2.怀疑go标准库中在 r.ParseForm 做了类似 iouitl.NoCloser 的操作，使我们读取body时，不是一次性的那种body，导致r.ParseForm()读取了后，依然能被json.NewDecoder()读取。但是，看了源码始终没有发现有类似的操作。

最终结果：
其实①病没有读取两次body，因返回的content-type为application/json时，1根本不会读取body的内容，这点从源码的ParseForm() --> parsePostForm() 可以看出只有content-type为application/x-www-form-urlencoded才会被解析。
所以上面给的第一个可能结果中【读取了两次body】是错误的。

由于2一定会读取body，1是根据content-type的值才会读取body，body正常情况下只能被读一次，所以必须要先做1再做2

 */
