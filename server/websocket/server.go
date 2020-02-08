package websocket

import (
	ws "github.com/gorilla/websocket"
	"github.com/sparkoo/payterm/server"
	"io"
	"log"
	"net/http"
)


var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
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

func NewServerWebsocket(addr string) server.Server {
	return &ServerWebsocket{addr: addr}
}

func (s *ServerWebsocket) Start() {
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal(err)
	}
}

// TODO: some proper close, stop and release release connections
func (s *ServerWebsocket) Close() {
}

func closeConnection(conn *ws.Conn) {
	if err := conn.Close(); err != nil {
		panic(err)
	}
}