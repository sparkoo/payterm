package websocket

import (
	ws "github.com/gorilla/websocket"
	"github.com/sparkoo/payterm/server"
	"io"
	"log"
	"net/http"
	"time"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ServerWebsocket struct {
	addr string
}

func (s *ServerWebsocket) AddHandler(addr string) io.ReadWriteCloser {
	h := newServerReadWriter()
	http.Handle(addr, h)
	return h
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

func pingTicker(writer io.Writer) {
	ticker := time.NewTicker(10 * time.Second)
	defer func() {
		ticker.Stop()
	}()
	for range ticker.C {
		_, _ = writer.Write([]byte("ping"))
	}
}
