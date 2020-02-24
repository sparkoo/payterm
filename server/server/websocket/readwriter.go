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
		log.Printf("connection created [%s]", request.URL)
	}
	defer closeConnection(conn)

	if readyErr := waitForReady(conn); readyErr != nil {
		log.Fatal(readyErr)
	}

	heartBeat := make(chan string)
	fail := make(chan error)

	go func() {
		// readloop
		for {
			if _, messageBytes, err := conn.ReadMessage(); err != nil {
				fail <- err
				return
			} else {
				message := string(messageBytes)
				//log.Printf("received [%s]", message)
				if message == pong {
					heartBeat <- message
				} else {
					if _, err := s.Read(messageBytes); err != nil {
						fail <- err
						return
					}
				}
			}
		}
	}()

	go func() {
		// writeloop
		for message := range s.write {
			log.Printf("about to write [%s]", message)
			if err := conn.WriteMessage(websocket.TextMessage, message); err != nil {
				fail <- err
				return
			}
		}
	}()

	go func() {
		// heartBeat
		if err := hartBeat(heartBeat, s.write); err != nil {
			log.Print("hearBeat failed: ", err)
			fail <- err
			return
		}
	}()

	go func() {
		hearth := time.NewTicker(1 * time.Second)
		for range hearth.C {
			s.write <- []byte("blabol")
		}
	}()

	for e := range fail {
		log.Print(e)
		//if err := conn.Close(); err != nil {
		//	log.Print("already closed")
		//	log.Print(err)
		//}
	}
}

func hartBeat(beat chan string, write chan []byte) error {
	hearth := time.NewTicker(1 * time.Second)
	for range hearth.C {
		write <- []byte(ping)
		message := <-beat
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
	log.Printf("waiting for ready ... ")
	_, messageBytes, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	message := string(messageBytes)
	log.Printf(message)

	if message == readyMessage {
		log.Println("ready received")
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
