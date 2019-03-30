package websocket

import (
	"bytes"
	"fmt"
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
	AddReadListener(addr string, reader io.Reader)
	Start()
	Close()
}

type ServerWebsocket struct {
	addr string
}

func (s *ServerWebsocket) AddWriteHandler(addr string) io.Writer {
	writer := NewServerWriter()
	http.Handle(addr, writer)
	return writer
}

func (s *ServerWebsocket) AddReadListener(addr string, reader io.Reader) {
	panic("implement me")
}

func NewServerWebsocket(addr string) Server {
	return &ServerWebsocket{addr: addr}
}

func handleReadRequest(writer http.ResponseWriter, request *http.Request) {
	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	c.PongHandler()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s, %v", message, mt)

		err = c.WriteMessage(websocket.TextMessage, bytes.NewBufferString("blob").Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (s *ServerWebsocket) Start() {
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (*ServerWebsocket) Close() {

}
