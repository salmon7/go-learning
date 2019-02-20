package main

import (
	"net/http"
	"log"
)

func main() {

	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("success"))
	}
	http.HandleFunc("test", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}

/*
压测http请求的方式有很多，其中一种是使用 ApacheBench

安装：
$ apt-get install apache2-utils

使用方式
ab -n 100000 -c 100 '127.0.0.1:8080/test'

表示 100并发，共10000个请求

 */
