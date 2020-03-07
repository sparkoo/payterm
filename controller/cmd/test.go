package main

import (
	"github.com/sparkoo/payterm/peripherals"
	"github.com/sparkoo/payterm/server/http"
	"log"
)

func main() {
	log.Println("starting terminal ...")

	server := http.NewServerHttp(":8080")
	testReader := server.AddReadListener("/testRead")
	//testWriter := server.AddWriteHandler("/testWrite")

	testPeripheralReader := peripherals.NewInputReader(testReader)
	//testPeripheralWriter := peripherals.NewOutputWriter(testWriter)

	go server.Start()

	for {
		//testPeripheralWriter.Write("hello")
		log.Print(testPeripheralReader.Read())
	}
}
