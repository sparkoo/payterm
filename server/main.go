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

	server := websocket.NewServerWebsocket(":8080")
	displayWriter := server.AddWriteHandler("/display")
	buzzerWriter := server.AddWriteHandler("/buzzer")
	keyboardReader := server.AddReadListener("/keyboard")
	cardreaderReader := server.AddReadListener("/cardreader")

	user1 := model.NewAccount(model.UserId("1"), "Jon Doe", 1000)
	log.Println(user1)

	users := make(map[model.UserId]*model.Account)
	users[user1.Id()] = user1

	var cardreader = peripherals.NewInputReader(cardreaderReader)
	var keyboard = peripherals.NewInputReader(keyboardReader)
	var display = peripherals.NewOutputWriter(displayWriter)
	var buzzer = peripherals.NewOutputWriter(buzzerWriter)

	terminal := term.NewTerm(server, users, keyboard, display, buzzer, cardreader)
	terminal.Start()
	defer terminal.Close()
}
