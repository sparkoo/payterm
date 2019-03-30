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
	var cardreader = peripherals.NewDummyCardReader()

	server := websocket.NewServerWebsocket(":8080")
	displayWriter := server.AddWriteHandler("/display")
	buzzerWriter := server.AddWriteHandler("/buzzer")

	go server.Start()

	var display = peripherals.NewDummyDisplay(displayWriter)
	var buzzer = peripherals.NewDummyBuzzer(buzzerWriter)
	for {
		display.Write("hello")
		buzzer.Beep()
		fmt.Println("loop")
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
