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
			fmt.Println("json.Marshal err: %v\n", err)
			continue
		}
		//fmt.Printf("%s\n", jdata)
	}

	if err := pprof.WriteHeapProfile(f); err != nil {
		fmt.Printf("could not write memory profile: %v\n", err)
		os.Exit(-1)
	}

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
			fmt.Println("json.Marshal err: %v\n", err)
			continue
		}
		//fmt.Printf("%s\n", jdata)
	}

	f.Close()
}

/*
zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof pprof_mem ./mem_profile.pf

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -alloc_space pprof_mem ./mem_profile.pf

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -alloc_object pprof_mem ./mem_profile.pf

zhang@debian-salmon-gb:~/Workspace/go/src/go-learning$ go tool pprof -inuse_objects pprof_mem ./mem_profile.pf


参考：
https://cizixs.com/2017/09/11/profiling-golang-program/
https://www.reddit.com/r/golang/comments/7ony5f/what_is_the_meaning_of_flat_and_cum_in_golang/


安装go-torch  go get github.com/uber/go-torch
 */
