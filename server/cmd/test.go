package main

import (
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/server/websocket"
	"log"
	"time"
)

func main() {
	log.Println("starting terminal ...")

	server := websocket.NewServerWebsocket(":8080")
	testReader := server.AddReadListener("/testRead")
	testWriter := server.AddWriteHandler("/testWrite")

	testPeripheralReader := peripherals.NewInputReader(testReader)
	testPeripheralWriter := peripherals.NewOutputWriter(testWriter)

	go server.Start()

	read := make(chan string)
	for {
		go func() {
			t := time.NewTicker(1 * time.Second)
			for range t.C {
				testPeripheralWriter.Write(time.Now().String())
			}
		}()
		go func() {
			read <- testPeripheralReader.Read()
		}()

		select {
		case card := <-read:
			log.Println("read: ", card)
		}
	}

}
