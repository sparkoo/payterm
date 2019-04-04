package peripherals

import (
	"io"
)

type KeyboardImpl struct {
	reader *peripheralReader
}

func (k *KeyboardImpl) Read() string {
	return k.reader.read()
}

func NewKeyboard(reader io.Reader) InputReader {
	return &KeyboardImpl{reader: newReader(reader)}
}
