package peripherals

import "fmt"

type Display interface {
	Write(string)
}

type Display2x16 struct {
}

func (*Display2x16) Write(message string) {
}

type DisplayDummy struct {
}

func (*DisplayDummy) Write(message string) {
	fmt.Printf("writing message: [%v]\n", message)
}

func NewDummyDisplay() Display {
	return &DisplayDummy{}
}

func NewDisplay() Display {
	return &Display2x16{}
}
