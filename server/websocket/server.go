package websocket

import (
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server interface {
	AddWriteHandler(addr string) io.Writer
	AddReadListener(addr string) io.Reader
	Start()
	Close()
}

type ServerWebsocket struct {
	addr string
}

func (s *ServerWebsocket) AddWriteHandler(addr string) io.Writer {
	writer := newServerWriter()
	http.Handle(addr, writer)
	return writer
}

func (s *ServerWebsocket) AddReadListener(addr string) io.Reader {
	r := newServerReader()
	http.Handle(addr, r)
	return r
}

func NewServerWebsocket(addr string) Server {
	return &ServerWebsocket{addr: addr}
}

func (s *ServerWebsocket) Start() {
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (*ServerWebsocket) Close() {

}

func closeConnection(conn *websocket.Conn) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
}
