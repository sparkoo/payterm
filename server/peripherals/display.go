package peripherals

import (
	"io"
)

type Display interface {
	Write(string)
}

type DisplayImpl struct {
	writer *peripheralWriter
}

func (d *DisplayImpl) Write(message string) {
	d.writer.write(message)
}

func NewDisplay(writer io.Writer) Display {
	return &DisplayImpl{writer: newWriter(writer)}
}
