package main

import (
	"fmt"
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/websocket"
	"time"
)

func main() {

	fmt.Println("starting terminal ...")
	var cardreader = peripherals.NewDummyCardReader()

	server := websocket.NewServerWebsocket(":8080")
	displayWriter := server.AddWriteHandler("/display")
	buzzerWriter := server.AddWriteHandler("/buzzer")
	keyboardReader := server.AddReadListener("/keyboard")

	go server.Start()

	var display = peripherals.NewDummyDisplay(displayWriter)
	var buzzer = peripherals.NewDummyBuzzer(buzzerWriter)
	var keyboard = peripherals.NewDummyKeyboard(keyboardReader)

	for {
		key := keyboard.Read()
		fmt.Printf("pressed [%s]\n", key)
		display.Write(key)
		//buzzer.Beep()
		//fmt.Println("loop")
		//time.Sleep(1 * time.Second)
	}
	buzzer.Beep()
	fmt.Printf("keyboard read: [%v]\n", keyboard.Read())
	fmt.Printf("cardreader read: [%v]\n", cardreader.Read())

	for {
		time.Sleep(1 * time.Second)
	}

	fmt.Println("ende")
}
