package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const readyMessage = "ready"
const ping = "ping"
const pong = "pong"

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

	if readyErr := waitForReady(conn); readyErr != nil {
		log.Fatal(readyErr)
	}

	go func() {
		if err := hartBeat(conn); err != nil {
			log.Print(err)
			closeConnection(conn)
		}
	}()

	go writeLoop(conn, s)
	go readLoop(conn, s)
	for {
		time.Sleep(1*time.Second)
	}
}

func hartBeat(conn *websocket.Conn) error {
	ping := []byte(ping)

	hearth := time.NewTicker(1 * time.Second)
	for range hearth.C {
		if err := conn.WriteMessage(websocket.TextMessage, ping); err != nil {
			return err
		}
		if messageType, messageBytes, err := conn.ReadMessage(); err != nil {
			return err
		} else {
			message := string(messageBytes)
			log.Printf("recv: %s, %v", message, messageType)
			if message != pong {
				return &invalidMessage{
					message:  message,
					expected: pong,
				}
			}
		}
	}
	return nil
}

func waitForReady(conn *websocket.Conn) error {
	messageType, messageBytes, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	message := string(messageBytes)
	log.Printf("recv: %s, %v", message, messageType)

	if message == readyMessage {
		return nil
	} else {
		return &invalidMessage{message: message, expected: readyMessage}
	}
}

func readLoop(conn *websocket.Conn, s *serverReadWriter) {
}

func writeLoop(conn *websocket.Conn, messageBus *serverReadWriter) {
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

type invalidMessage struct {
	message  string
	expected string
}

func (err *invalidMessage) Error() string {
	return fmt.Sprintf("expected [%s] but got [%s]", err.expected, err.message)
}
