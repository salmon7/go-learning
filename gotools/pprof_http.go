package main

import (
	_ "net/http/pprof"
	"net/http"
	"fmt"
	"os"
)

func main() {
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("http start error")
		os.Exit(-1)
	}
}

/*
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go build gotools/pprof_http.go
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ ./pprof_http &

安装go-torch(go-torch 也能从本地生成火焰图，不一定要访问url)
go get github.com/uber/go-torch

在项目根目录运行
git clone https://github.com/brendangregg/FlameGraph.git


参考：
https://cizixs.com/2017/09/11/profiling-golang-program/
https://www.reddit.com/r/golang/comments/7ony5f/what_is_the_meaning_of_flat_and_cum_in_golang/

pprof
https://github.com/google/pprof/blob/master/doc/README.md

Profiling Go Programs
https://blog.golang.org/profiling-go-programs

graphviz
https://graphviz.gitlab.io/about/

 */