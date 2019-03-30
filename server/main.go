package main

import (
	"fmt"
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/websocket"
	"time"
)

func main() {

	fmt.Println("starting terminal ...")
	var keyboard = peripherals.NewDummyKeyboard()
	var buzzer = peripherals.NewDummyBuzzer()
	var cardreader = peripherals.NewDummyCardReader()

	server := websocket.NewServerWebsocket(":8080")
	displayWriter := server.AddWriteHandler("/display")

	go server.Start()

	var display = peripherals.NewDummyDisplay(displayWriter)
	for {
		display.Write("hello")
		time.Sleep(1 * time.Second)
	}
	buzzer.Beep()
	fmt.Printf("keyboard read: [%v]\n", keyboard.Read())
	fmt.Printf("cardreader read: [%v]\n", cardreader.Read())

	for {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ende")
}
