package main

import (
	"fmt"
	"github.com/sparkoo/payterm/model"
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/term"
	"github.com/sparkoo/payterm/websocket"
	"log"
)

func main() {

	fmt.Println("starting terminal ...")
	var cardreader = peripherals.NewCardReader()

	server := websocket.NewServerWebsocket(":8080")
	displayWriter := server.AddWriteHandler("/display")
	buzzerWriter := server.AddWriteHandler("/buzzer")
	keyboardReader := server.AddReadListener("/keyboard")

	user1 := model.NewAccount("Jon Doe", 1000)
	log.Println(user1)

	users := make(map[model.UserId]*model.Account)
	users[user1.Id()] = user1


	var display = peripherals.NewDisplay(displayWriter)
	var buzzer = peripherals.NewBuzzer(buzzerWriter)
	var keyboard = peripherals.NewKeyboard(keyboardReader)

	terminal := term.NewTerm(server, users, &keyboard, &display, &buzzer, &cardreader)
	terminal.Start()
	defer terminal.Close()
}
