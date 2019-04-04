package peripherals

import "io"

type CardReaderImpl struct {
	reader *peripheralReader
}

func (cr *CardReaderImpl) Read() string {
	return cr.reader.read()
}

func NewCardReader(reader io.Reader) InputReader {
	return &CardReaderImpl{reader: newReader(reader)}
}
