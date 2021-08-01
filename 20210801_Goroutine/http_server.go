package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello,http server"))
}

func StartHttpServer(server *http.Server) error {
	http.HandleFunc("/", handler)
	fmt.Println("StartHttpServer success")
	err := server.ListenAndServe()
	return err
}

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	group, errCtx := errgroup.WithContext(ctx)

	server := &http.Server{Addr: "127.0.0.1:7777"}

	//开启和关闭httpserver
	group.Go(func() error {
		return StartHttpServer(server)
	})

	group.Go(func() error {
		<-errCtx.Done()
		fmt.Println("Stop Http Server")
		return server.Shutdown(errCtx)
	})

	//监听linux signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch)

	//错误处理goroutine
	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done():
				return errCtx.Err()
			case <-ch:
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("error: ", err)
	}
}
