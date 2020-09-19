package main

import (
	"net/http"
	"os"
)

type handler struct{}

func main() {
	h := new(handler)
	err := http.ListenAndServe(":8888", h)
	if err != nil {
		return
	}
}

func (h *handler) ServeHTTP(rsp http.ResponseWriter, req *http.Request) {
	url := *req.URL
	q := url.Query()
	src := q["src"][0]
	go func() {
		f, _ := os.OpenFile("/tmp/out.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0777)
		f.Write([]byte(src))
		f.Write([]byte("\n"))
	}()
	rsp.WriteHeader(200)
}
