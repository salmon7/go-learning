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

1.多了几个endpoint
使用net/http/pprof后
你的 HTTP 服务都会多出 /debug/pprof endpoint，访问它会得到类似下面的内容：

/debug/pprof/

profiles:
0    block
62    goroutine
444    heap
30    threadcreate

full goroutine stack dump
这个路径下还有几个子页面：

/debug/pprof/profile：访问这个链接会自动进行 CPU profiling，持续 30s，并生成一个文件供下载
/debug/pprof/heap： Memory Profiling 的路径，访问这个链接会得到一个内存 Profiling 结果的文件
/debug/pprof/block：block Profiling 的路径
/debug/pprof/goroutines：运行的 goroutines 列表，以及调用关系


2.生成profile
默认会存到 /home/zhang/pprof/ 目录下
cpu profile:
go tool pprof ./pprof_http http://localhost:8080/debug/pprof/profile

mem profile:
go tool pprof ./pprof_http http://localhost:8080/debug/pprof/heap
这里也可以添加具体的指标参数
-inuse_space 默认，内存的使用空间大小
-inuse_objects 使用的对象数
-alloc-space 分配的内存空间大小
-alloc-objects 分配的对象数

3.生成火焰图
3.1安装go-torch(go-torch 访问url的形式，也能从本地生成火焰图)
go get github.com/uber/go-torch

在项目根目录运行
git clone https://github.com/brendangregg/FlameGraph.git

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go-torch -u http://localhost:8080 -f pprof_http.svg

go-torch ./pprof_http /home/zhang/pprof/pprof.pprof_http.samples.cpu.009.pb.gz -f http_profile.svg

3.2 使用go tool pprof -http=":8081"启动web客户端，从web客户端中查看火焰图以及函数调用图，top，list等，其中火焰图需要go 1.11以上版本
go tool pprof -http=":8081" ./pprof_http http://localhost:8080/debug/pprof/profile

2.3 直接使用pprof
# Get the pprof tool directly
$ go get -u github.com/google/pprof

$ pprof -http=":8081" ./pprof_http http://localhost:8080/debug/pprof/profile

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