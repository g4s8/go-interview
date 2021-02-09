package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	out = "/tmp/out.txt"
)

// http request handler
type handler struct {
	cnt       int32       // request counter
	writeChan chan string // channel for `src` query param from URLs
}

func main() {
	h := new(handler)
	h.writeToOut()
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *handler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	// parse `src` query param from URL
	url := *req.URL
	q := url.Query()
	src := q["src"]
	if src == nil {
		rsp.WriteHeader(http.StatusBadRequest)
		return
	}
	go func() { h.writeChan <- src[0] }()
}

// write request counter and query param form handler.writeChan to file
func (h *handler) writeToOut() {

	sigChan := make(chan os.Signal, 1)
	h.writeChan = make(chan string)

	signal.Notify(sigChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		//Create file if it does not exist, or truncate if exists
		f, err := os.Create(out)
		if err != nil {
			log.Fatal(err)
		}
		for {
			select {
			case s := <-h.writeChan:
				_, err := f.Write([]byte(fmt.Sprintf("%d: %v \n", h.cnt, s)))
				if err != nil {
					log.Fatal(err)
				}
				h.cnt++
			case sig := <-sigChan:
				//Close file and exit
				err := f.Close()
				if err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				fmt.Println(sig)
				os.Exit(0)
			}
		}
	}()
}
