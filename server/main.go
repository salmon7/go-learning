package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var (
	server   *Server
	graceful = flag.Bool("graceful", false, "listen on fd open 3 (internal use only)")
)

// NewMux return new Mux
func NewMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello world"))
	})
	return mux
}

func main() {
	flag.Parse()
	mux := NewMux()
	server = &Server{
		Server: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
		Pid:           os.Getpid(),
		MaxReloadTime: 10,
		sig:           make(chan os.Signal),
		graceful:      *graceful,
	}
	// server.Server.SetKeepAlivesEnabled(false)

	if err := server.Serve(); err != nil {
		log.Fatalf("%v, pid=%d", err, os.Getpid())
	}
	//
	// f := os.NewFile(3, "")
	// listener, err = net.FileListener(f)
	//
	// cmd := exec.Command(os.Args[0])
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	// // put socket FD at the first entry
	// cmd.ExtraFiles = []*os.File{f}

}

/*
测试方法：

1.终端一
$ go build -o main server.go main.go
$ ./main

输出：
2019/07/09 20:56:42 listening, pid=8238
2019/07/09 20:59:18 receive SIGUSR2 user defined signal 2, pid: 8238
2019/07/09 20:59:18 no longer serve, pid=8238
2019/07/09 20:59:18 listening, pid=8790
2019/07/09 20:59:18 gracefully reload success, pid: 8238, child's pid: 8790

2019/07/09 20:59:30 receive SIGUSR2 user defined signal 2, pid: 8790
2019/07/09 20:59:30 no longer serve, pid=8790
2019/07/09 20:59:30 listening, pid=8860
2019/07/09 20:59:30 gracefully reload success, pid: 8790, child's pid: 8860


2.终端二
$ ab -n 1000000 -c 10 "http://127.0.0.1:8080/hello"

3.终端三
$ ps aux | grep main | grep -v grep
zhang     8790  138  0.1 446048  9460 pts/1    Sl   20:59   0:08 ./main -graceful
$ kill -s SIGUSR2 8238

$ ps aux | grep main | grep -v grep
zhang     8860  5.9  0.1 446112  6584 pts/1    Sl   20:59   0:33 ./main -graceful
$ kill -s SIGUSR2 8790

用这个也行:$ pid=$(ps aux | grep main | grep -v grep | awk '{print $2}'); kill -s SIGUSR2 $pid

关注终端二这行的输出,如果为0则表示没有丢失请求：
Failed requests:        0


注意配置：
在程序重启时，ab偶尔会报 apr_socket_recv: Connection reset by peer (104)错误，可以参考以下配置，减少报错的概率(https://www.cnblogs.com/felixzh/p/8295471.html)
$ sudo vim /etc/sysctl.conf
net.ipv4.tcp_syncookies = 0
$ sudo sysctl -p

*/
