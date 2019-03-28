package main

import (
	"fmt"
	"github.com/sparkoo/payterm/peripherals"
)

func main() {
	fmt.Println("starting terminal ...")
	var display = peripherals.NewDummyDisplay()
	display.Write("hello")
	fmt.Println("ende")
}
