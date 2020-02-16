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

	read := make(chan string)
	write := make(chan string)
	heartBeat := make(chan string)
	fail := make(chan error)

	go func() {
		for {
			if _, messageBytes, err := conn.ReadMessage(); err != nil {
				fail <- err
			} else {
				message := string(messageBytes)
				read <- message
				heartBeat <- message
			}
		}
	}()

	go func(write chan string) {
		for message := range write{
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				fail <- err
			}
		}
	}(write)

	go func() {
		if err := hartBeat(heartBeat, write); err != nil {
			log.Print("hearBeat failed: ", err)
			fail <- err
		}
	}()

	go func() {
		if err := writeLoop(conn, s); err != nil {
			log.Print("writeLoop failed: ", err)
			fail <- err
		}
	}()

	go func() {
		if err := readLoop(read, s); err != nil {
			log.Print("readLoop failed: ", err)
			fail <- err
		}
	}()

	for e := range fail {
		log.Print(e)
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}
}

func hartBeat(beat chan string, write chan string) error {
	hearth := time.NewTicker(1 * time.Second)
	for range hearth.C {
		write <- ping
		message := <- beat
		if message != pong {
			return &invalidMessage{
				message:  message,
				expected: pong,
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

func readLoop(read chan string, s *serverReadWriter) error {
	for message := range read {
		if _, err := s.Read([]byte(message)); err != nil {
			return err
		}
	}
	return nil
}

func writeLoop(conn *websocket.Conn, messageBus *serverReadWriter) error {
	return nil
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
