package websocket

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type serverReader struct {
	messageBus chan []byte
}

func (serverReader *serverReader) Read(messageBytes []byte) (n int, err error) {
	copy(messageBytes, <-serverReader.messageBus)
	return len(messageBytes), nil
}

func (serverReader *serverReader) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	} else {
		log.Println("connection created", request.URL)
	}
	defer closeConnection(conn)

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			break
		}
		log.Printf("recv: %s, %v", message, mt)
		serverReader.messageBus <- message

		err = conn.WriteMessage(websocket.TextMessage, bytes.NewBufferString("blob").Bytes())
		if err != nil {
			fmt.Println(err)
		}
	}
}

func newServerReader() *serverReader {
	return &serverReader{messageBus: make(chan []byte)}
}
