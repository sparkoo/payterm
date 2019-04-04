package peripherals

import (
	"io"
)

type BuzzerImpl struct {
	writer *peripheralWriter
}

func (b *BuzzerImpl) Write(msg string) {
	b.writer.write("beep")
}

func NewBuzzer(writer io.Writer) OutputWriter {
	return &BuzzerImpl{writer: newWriter(writer)}
}
