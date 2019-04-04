package peripherals

import (
	"io"
)

type DisplayImpl struct {
	writer *peripheralWriter
}

func (d *DisplayImpl) Write(message string) {
	d.writer.write(message)
}

func NewDisplay(writer io.Writer) OutputWriter {
	return &DisplayImpl{writer: newWriter(writer)}
}
