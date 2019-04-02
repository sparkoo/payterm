package peripherals

import (
	"io"
)

type Buzzer interface {
	Beep()
}

type BuzzerImpl struct {
	writer *peripheralWriter
}

func (b *BuzzerImpl) Beep() {
	b.writer.write("beep")
}

func NewBuzzer(writer io.Writer) Buzzer {
	return &BuzzerImpl{writer: newWriter(writer)}
}
