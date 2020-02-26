package http

import (
	"github.com/sparkoo/payterm/server"
	"io"
	"log"
	"net/http"
)

type ServerHttp struct {
	addr string
}

func (s *ServerHttp) AddWriteHandler(addr string) io.Writer {
	return newHttpServerWriter(addr)
}

func (s *ServerHttp) AddReadListener(addr string) io.Reader {
	return newHttpServerReader(addr)
}

func (s *ServerHttp) Start() {
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (s *ServerHttp) Close() {
	panic("implement me")
}

func NewServerHttp(addr string) server.Server {
	return &ServerHttp{addr: addr}
}
