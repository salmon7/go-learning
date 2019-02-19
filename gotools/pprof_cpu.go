package main

import (
	"fmt"
	"runtime/pprof"
	"os"
	"encoding/json"
)

func main() {
	f, err := os.Create("cpu_profile.pf")
	if err != nil {
		fmt.Printf("os.Creat err: %v\n", err)
		os.Exit(-1)
	}
	defer f.Close()

	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	for i := 0; i < 10000000; i++ {
		Person := struct {
			Name string
			Age  int
		}{
			"123",
			i,
		}
		_, err := json.Marshal(Person)
		if err != nil {
			fmt.Printf("json.Marshal err: %v\n", err)
			continue
		}
		//fmt.Printf("%s\n", jdata)
	}

}

/*
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go build gotools/pprof_cpu.go
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ ./pprof_cpu
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof pprof_cpu cpu_profile.pf
File: pprof_1
Type: cpu
Time: Feb 1, 2019 at 11:04am (CST)
Duration: 7.24s, Total samples = 6.82s (94.22%)
Entering interactive mode (type "help" for commands, "o" for options)
(pprof) top20
Showing nodes accounting for 4260ms, 62.46% of 6820ms total
Dropped 46 nodes (cum <= 34.10ms)
Showing top 20 nodes out of 83
      flat  flat%   sum%        cum   cum%
     510ms  7.48%  7.48%     2770ms 40.62%  encoding/json.(*structEncoder).encode
     430ms  6.30% 13.78%      510ms  7.48%  bytes.(*Buffer).WriteByte
     410ms  6.01% 19.79%     1170ms 17.16%  runtime.mallocgc
     290ms  4.25% 24.05%      720ms 10.56%  encoding/json.(*encodeState).string
     280ms  4.11% 28.15%      640ms  9.38%  encoding/json.intEncoder
     220ms  3.23% 31.38%     5720ms 83.87%  encoding/json.Marshal
     220ms  3.23% 34.60%      220ms  3.23%  runtime.heapBitsSetType
     210ms  3.08% 37.68%      210ms  3.08%  runtime.memequal64
     160ms  2.35% 40.03%     4760ms 69.79%  encoding/json.(*encodeState).marshal
     160ms  2.35% 42.38%      160ms  2.35%  runtime.aeshash64
     160ms  2.35% 44.72%      760ms 11.14%  runtime.mapaccess2
     160ms  2.35% 47.07%      160ms  2.35%  runtime.memmove
     150ms  2.20% 49.27%      150ms  2.20%  runtime.procyield
     140ms  2.05% 51.32%      140ms  2.05%  runtime.memclrNoHeapPointers
     140ms  2.05% 53.37%      300ms  4.40%  runtime.nilinterhash
     130ms  1.91% 55.28%      200ms  2.93%  encoding/json.fieldByIndex
     130ms  1.91% 57.18%      160ms  2.35%  strconv.formatBits
     120ms  1.76% 58.94%      200ms  2.93%  bytes.(*Buffer).WriteString
     120ms  1.76% 60.70%     2890ms 42.38%  encoding/json.(*structEncoder).(encoding/json.encode)-fm
     120ms  1.76% 62.46%      530ms  7.77%  encoding/json.stringEncoder

如何查看数据：
第一列表示该函数(不包括子函数)的CPU运行时间
第二列表示该函数(不包括子函数)的CPU运行时间占CPU百分比，如第一行约等于 510/6820=7.48%
第三列表示从上往下所有函数累加使用CPU的比例，与具体函数没有关系，sumN = flatN + sum(N-1)
第四列表示该函数及其子函数的CPU运行时间
第五列表示该函数及其子函数的CPU运行时间占CPU百分比
第六列表示函数的名字

1.在pprof交互命令行中，输入web命令查看函数调用图
(pprof) web

2.生成火焰图
2.1 使用go-torch时候生成火焰图
安装go-torch(go-torch 也能从本地生成火焰图，不一定要访问url)
go get github.com/uber/go-torch

在项目根目录运行
git clone https://github.com/brendangregg/FlameGraph.git

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go-torch ./pprof_cpu cpu_profile.pf -f cpu_profile.svg

2.2 使用go tool pprof -http=":8081"启动web客户端，从web客户端中查看火焰图以及函数调用图，top，list等，其中火焰图需要go 1.11以上版本
# This will listen on :8081 and open a browser.
# Change :8081 to a port of your choice.
go tool pprof -http=":8081" ./pprof_cpu cpu_profile.pf

2.3 直接使用pprof
# Get the pprof tool directly
$ go get -u github.com/google/pprof

$ pprof -http=":8081" ./pprof_cpu cpu_profile.pf


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