package peripherals

import (
	"fmt"
	"io"
)

type Display interface {
	Write(string)
}

type Display2x16 struct {
}

func (*Display2x16) Write(message string) {
}

type DisplayDummy struct {
	writer io.Writer
}

func (d *DisplayDummy) Write(message string) {
	fmt.Printf("writing message: [%v]\n", message)
	if _, err := d.writer.Write([]byte(message)); err != nil {
		fmt.Printf("error writing [%v]", err)
	}
}

func NewDummyDisplay(writer io.Writer) Display {
	return &DisplayDummy{writer: writer}
}

func NewDisplay() Display {
	return &Display2x16{}
}
