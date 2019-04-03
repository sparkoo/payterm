package peripherals

import (
	"io"
)

type Keyboard interface {
	Read() string
}

type KeyboardImpl struct {
	reader *peripheralReader
}

func (k *KeyboardImpl) Read() string {
	return k.reader.read()
}

func NewKeyboard(reader io.Reader) Keyboard {
	return &KeyboardImpl{reader: newReader(reader)}
}
