package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

type httpServerReader struct {
	read chan []byte
}

func (h *httpServerReader) Read(p []byte) (n int, err error) {
	copy(p, <-h.read)
	return len(p), nil
}

func newHttpServerReader(addr string) *httpServerReader {
	read := &httpServerReader{read: make(chan []byte)}
	http.HandleFunc(addr, func(w http.ResponseWriter, r *http.Request) {
		if bytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			read.read <- bytes
		}
	})
	return read
}