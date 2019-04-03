package peripherals

import "io"

type CardReader interface {
	Read() string
}

type CardReaderImpl struct {
	reader *peripheralReader
}

func (cr *CardReaderImpl) Read() string {
	return cr.reader.read()
}

func NewCardReader(reader io.Reader) CardReader {
	return &CardReaderImpl{reader: newReader(reader)}
}
