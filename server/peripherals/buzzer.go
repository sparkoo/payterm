package peripherals

import (
	"io"
)

type Buzzer interface {
	Beep()
}

type BuzzerDummy struct {
	writer *peripheralWriter
}

func (b *BuzzerDummy) Beep() {
	b.writer.write("beep")
}

func NewDummyBuzzer(writer io.Writer) Buzzer {
	return &BuzzerDummy{writer: newWriter(writer)}
}
