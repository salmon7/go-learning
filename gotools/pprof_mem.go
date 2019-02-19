package main

import (
	"fmt"
	"os"
	"encoding/json"
	"runtime/pprof"
)

func main() {
	f, err := os.Create("mem_profile.pf")
	if err != nil {
		fmt.Printf("os.Creat err: %v\n", err)
		os.Exit(-1)
	}

	for i := 0; i < 1000; i++ {
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

	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("could not write memory profile: %v\n", err)
		os.Exit(-1)
	}
	f.Close()
}

/*

-inuse_space 默认，内存的使用空间大小
-inuse_objects 使用的对象数
-alloc-space 分配的内存空间大小
-alloc-objects 分配的对象数

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof pprof_mem ./mem_profile.pf
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -inuse_objects pprof_mem ./mem_profile.pf
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -alloc_space pprof_mem ./mem_profile.pf
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -alloc_objects pprof_mem ./mem_profile.pf


火焰图生成方式参见 pprof_cpu.go

参考：
https://cizixs.com/2017/09/11/profiling-golang-program/
https://www.reddit.com/r/golang/comments/7ony5f/what_is_the_meaning_of_flat_and_cum_in_golang/

 */
