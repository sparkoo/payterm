package main

import (
	"fmt"
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/websocket"
	"time"
)

func main() {

	fmt.Println("starting terminal ...")
	var display = peripherals.NewDummyDisplay()
	var keyboard = peripherals.NewDummyKeyboard()
	var buzzer = peripherals.NewDummyBuzzer()
	var cardreader = peripherals.NewDummyCardReader()

	server := websocket.NewServerWebsocket(":8080")
	go server.Start()

	display.Write("hello")
	buzzer.Beep()
	fmt.Printf("keyboard read: [%v]\n", keyboard.Read())
	fmt.Printf("cardreader read: [%v]\n", cardreader.Read())

	for {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ende")
}
