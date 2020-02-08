package tcp

import (
	"fmt"
	"github.com/sparkoo/payterm/server"
	"io"
	"log"
	"net"
)

type ServerTcp struct {
	addr string

}

func NewServerTcp(addr string) server.Server {
	return &ServerTcp{addr: addr}
}

func (s *ServerTcp) AddWriteHandler(addr string) io.Writer {
	panic("implement me")
}

func (s *ServerTcp) AddReadListener(addr string) io.Reader {
	panic("implement me")
}

func (s *ServerTcp) Start() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	fmt.Printf("%+v", conn)
}

func (s *ServerTcp) Close() {
	panic("implement me")
}
