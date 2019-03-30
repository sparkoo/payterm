package websocket

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type serverWriter struct {
	messages chan string
}

func NewServerWriter() *serverWriter {
	return &serverWriter{messages: make(chan string)}
}

func (s *serverWriter) Write(p []byte) (n int, err error) {
	s.messages <- string(p)
	return len(p), nil
}

func (s *serverWriter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	c, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	go writeloop(c, s)
}

func writeloop(c *websocket.Conn, s *serverWriter) {
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			fmt.Println("connection closed")
			if err := c.Close(); err != nil {
				fmt.Println("error closing connection", err)
			}
			break
		}
		log.Printf("recv: %s, %v", message, mt)

		if string(message) == "ready" {
			writeMessage := <-s.messages
			log.Println("writing message ", writeMessage)
			err = c.WriteMessage(websocket.TextMessage, bytes.NewBufferString(writeMessage).Bytes())
			if err != nil {
				fmt.Println("connection closed")
				if err := c.Close(); err != nil {
					fmt.Println("error closing connection", err)
				}
				break
			}
		}
	}
}
