package http

import (
	"io/ioutil"
	"log"
	"net/http"
)

type httpServerWriter struct {
	write chan []byte
}

func (h *httpServerWriter) Write(p []byte) (n int, err error) {
	h.write <- p
	return len(p), nil
}

func newHttpServerWriter(addr string) *httpServerWriter {
	write := &httpServerWriter{write: make(chan []byte, 1024)}
	http.HandleFunc(addr, func(w http.ResponseWriter, r *http.Request) {
		if bytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			ready := string(bytes)
			if ready == "ready" {
				if _, err := w.Write(<-write.write); err != nil {
					log.Fatal(err)
				}
			}
		}
	})
	return write
}
