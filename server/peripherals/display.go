package peripherals

import (
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
	writer *peripheralWriter
}

func (d *DisplayDummy) Write(message string) {
	d.writer.write(message)
}

func NewDummyDisplay(writer io.Writer) Display {
	return &DisplayDummy{writer: newWriter(writer)}
}

func NewDisplay() Display {
	return &Display2x16{}
}
