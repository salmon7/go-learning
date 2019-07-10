package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// Server elegant stop server
type Server struct {
	Server        *http.Server
	Pid           int
	MaxReloadTime int

	child    *exec.Cmd
	listen   net.Listener
	sig      chan os.Signal
	graceful bool
}

// Serve serve
func (s *Server) Serve() error {

	var err error
	if s.graceful {
		// 用reload方式启动子进程时，子进程会进入这里，
		// 获取父进程传递过来的文件句柄
		f := os.NewFile(3, "")

		// 复制父进程传递过来的文件句柄
		// Closing ln does not affect f, and closing f does not affect ln.
		s.listen, err = net.FileListener(f)
		if err != nil {
			log.Fatalf("reload err: %v, pid=%d", err, s.Pid)
		}

		// 创建 reload.${ppid} 标志子进程启动
		reloadFileName := fmt.Sprintf("reload.%d", os.Getppid())
		log.Printf("create reload file: %s, pid=%d", reloadFileName, s.Pid)
		reloadFile, err := os.Create(reloadFileName)
		if err != nil {
			log.Fatalf("reload error: %v, pid=%d", err, s.Pid)
		}
		reloadFile.Close()
	} else {
		// 正常启动进程时，将监听 tcp Addr
		s.listen, err = net.Listen("tcp", s.Server.Addr)
		if err != nil {
			log.Fatalf("start err: %v, pid=%d", err, os.Getpid())
		}
	}

	log.Printf("listening, pid=%d", s.Pid)
	signal.Notify(s.sig, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR2)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		if err := s.Server.Serve(s.listen); err != nil {
			if err == http.ErrServerClosed {
				log.Printf("no longer serve, pid=%d", s.Pid)
			} else {
				log.Printf("Serve err: %v, pid=%d", err, s.Pid)
			}
		}
		wg.Done()
	}()

	var errCh = make(chan error)
	reloadFileName := fmt.Sprintf("reload.%d", s.Pid)
	for {
		c := <-s.sig
		switch c {
		case syscall.SIGINT, syscall.SIGTERM:
			log.Printf("receive SIGINT/SIGTERM %s, pid: %d", c, os.Getpid())
			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err := server.Server.Shutdown(ctx)
			if err != nil {
				log.Printf("Shutdown err :%v , pid: %d, child's pid: %d", err, s.Pid, s.child.Process.Pid)
			}

			log.Printf("gracefully shutdown, pid: %d", os.Getpid())
			wg.Wait()
			return nil
		case syscall.SIGUSR2:
			log.Printf("receive SIGUSR2 %s, pid: %d", c, os.Getpid())
			err := s.Reload()
			if err != nil {
				log.Printf("gracefully reload error: %v, pid: %d", err, os.Getpid())
				break
			}

			// shutdown前，通过reload.pid判断子进程是否启动成功，只有子进程提供服务了才shutdown
			go func() {
				for i := 0; i < s.MaxReloadTime; i++ {
					_, err = os.Stat(reloadFileName)
					if err == nil {
						e := os.Remove(reloadFileName)
						if e != nil {
							log.Printf("clear reload file err: %s, pid=%d", e, s.Pid)
						}
						// new process has started succ
						errCh <- nil
						return
					}
					time.Sleep(time.Second)
				}
				// avoid to block this goroutine
				select {
				case errCh <- fmt.Errorf(
					"Wait for reload ready timeout: %vs", s.MaxReloadTime):
				case <-time.After(time.Second):
					log.Printf("send error to channel timeout, pid=%d", s.Pid)
				}

			}()

			err = <-errCh
			if err != nil {
				os.Remove(reloadFileName)
				log.Printf("try to reload error: %s, pid=%d", err, s.Pid)
				return err
			}

			ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
			err = server.Server.Shutdown(ctx)
			if err != nil {
				log.Printf("Shutdown err: %v , pid=%d, child's pid=%d", err, s.Pid, s.child.Process.Pid)
			}

			log.Printf("gracefully reload success, pid=%d, child's pid=%d", s.Pid, s.child.Process.Pid)
			wg.Wait()
			return nil
		}
	}
	wg.Wait()
	return nil
}

// Reload start a child with socket fd
func (s *Server) Reload() error {
	tl, ok := s.listen.(*net.TCPListener)
	if !ok {
		return errors.New("listener is not tcp listener")
	}

	f, err := tl.File()
	if err != nil {
		return err
	}

	args := []string{"-graceful"}
	s.child = exec.Command(os.Args[0], args...)
	s.child.Stdout = os.Stdout
	s.child.Stderr = os.Stderr
	// put socket FD at the first entry
	s.child.ExtraFiles = []*os.File{f}
	return s.child.Start()
}
