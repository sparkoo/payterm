package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type serverReadWriter struct {
	write chan []byte
	read  chan []byte
}

func (s *serverReadWriter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	} else {
		log.Println("connection created", request.URL)
	}
	defer closeConnection(conn)

	go writeLoop(conn, s)
	go readLoop(conn, s)
}

func readLoop(conn *websocket.Conn, s *serverReadWriter) {
	panic("implement me!")
}

func writeLoop(conn *websocket.Conn, messageBus *serverReadWriter) {
	panic("implement me!")
}

func (s *serverReadWriter) Read(readMessage []byte) (n int, err error) {
	copy(readMessage, <-s.read)
	return len(readMessage), nil
}

func (s *serverReadWriter) Write(writeMessage []byte) (n int, err error) {
	s.write <- writeMessage
	return len(writeMessage), nil
}

func (s *serverReadWriter) Close() error {
	panic("implement me")
}

func newServerReadWriter() *serverReadWriter {
	return &serverReadWriter{write: make(chan []byte), read: make(chan []byte)}
}
