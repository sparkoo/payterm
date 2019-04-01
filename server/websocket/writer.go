package websocket

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

const ready = "ready"

type serverWriter struct {
	messages chan []byte
}

func newServerWriter() *serverWriter {
	return &serverWriter{messages: make(chan []byte)}
}

func (serverWriter *serverWriter) Write(p []byte) (n int, err error) {
	serverWriter.messages <- p
	return len(p), nil
}

func (serverWriter *serverWriter) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	conn, err := upgrader.Upgrade(writer, request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	if err := writeloop(conn, serverWriter); err != nil {
		fmt.Println("connection closed")
		closeConnection(conn)
	}
}

func writeloop(conn *websocket.Conn, messageBus *serverWriter) error {
	for {
		if err := waitForReady(conn); err == nil {
			writeMessage := <-messageBus.messages
			log.Println("writing message ", writeMessage)
			if err := conn.WriteMessage(websocket.TextMessage, writeMessage);
				err != nil {
				return err
			}
		} else {
			log.Println(err)
		}
	}
}

func waitForReady(conn *websocket.Conn) error {
	mt, messageBytes, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("connection closed")
		if err := conn.Close(); err != nil {
			fmt.Println("error closing connection", err)
		}
		return err
	}
	log.Printf("recv: %s, %v", messageBytes, mt)
	msg := string(messageBytes)
	if msg == ready {
		return nil
	} else {
		return &invalidReadyMessage{message: msg}
	}
}

type invalidReadyMessage struct {
	message string
}

func (err *invalidReadyMessage) Error() string {
	return fmt.Sprintf("expected [ready] but got [%v]", err.message)
}
