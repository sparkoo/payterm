package websocket

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Server interface {
	Write(addr string, message string)
	AddReadListener(addr string, read func(message string))
	Start()
	Close()
}

type ServerWebsocket struct {
	addr string
}

func (s *ServerWebsocket) Write(addr string, message string) {
	panic("implement me")
}

func (s *ServerWebsocket) AddReadListener(addr string, read func(message string)) {
	panic("implement me")
}

func NewServerWebsocket(addr string) Server {
	return &ServerWebsocket{addr: addr}
}

func (*ServerWebsocket) init() {
	http.HandleFunc("/keyboard", handleReadRequest)
	http.HandleFunc("/cardreader", handleReadRequest)

	http.HandleFunc("/display", handleWriteRequest)
	http.HandleFunc("/buzzer", handleWriteRequest)
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

func handleWriteRequest(writer http.ResponseWriter, request *http.Request) {
	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("connection closed")
			break
		}
		log.Printf("recv: %s, %v", message, mt)

		time.Sleep(500 * time.Millisecond)

		err = c.WriteMessage(websocket.TextMessage, bytes.NewBufferString("blob").Bytes())
		if err != nil {
			fmt.Println("connection closed")
			break
		}
	}
}

func (s *ServerWebsocket) Start() {
	s.init()
	if err := http.ListenAndServe(s.addr, nil); err != nil {
		log.Fatal(err)
	}
}

func (*ServerWebsocket) Close() {

}
