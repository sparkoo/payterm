package main

import (
	"fmt"
	"github.com/sparkoo/payterm/peripherals"
)

func main() {
	fmt.Println("starting terminal ...")
	var display = peripherals.NewDummyDisplay()
	var keyboard = peripherals.NewDummyKeyboard()
	var buzzer = peripherals.NewDummyBuzzer()
	var cardreader = peripherals.NewDummyCardReader()

	display.Write("hello")
	buzzer.Beep()
	fmt.Printf("keyboard read: [%v]\n", keyboard.Read())
	fmt.Printf("cardreader read: [%v]\n", cardreader.Read())

	fmt.Println("ende")
}
