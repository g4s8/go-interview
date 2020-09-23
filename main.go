package main

import (
	"fmt"
	"net/http"
	"os"
)

const (
	out = "/tmp/out.txt"
)

// http request handler
type handler struct {
	// request counter
	cnt int32
}

func main() {
	// remove output file before starting
	os.Remove(out)

	h := new(handler)
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		return
	}
}

func (h handler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	// parse `src` query param from URL
	url := *req.URL
	q := url.Query()
	src := q["src"][0]
	go func() {
		// write request counter and query param to file
		f, _ := os.OpenFile(out, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		f.Write([]byte(fmt.Sprintf("%d: ", h.cnt)))
		h.cnt++
		f.Write([]byte(src))
		f.Write([]byte("\n"))
	}()
	// respond with 200 - OK
	rsp.WriteHeader(200)
}
