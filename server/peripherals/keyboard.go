package peripherals

import (
	"io"
	"log"
)

type Keyboard interface {
	Read() string
}

type KeyboardImpl struct {
	reader io.Reader
}

func (k *KeyboardImpl) Read() string {
	keyBytes := make([]byte, 1)
	_, err := k.reader.Read(keyBytes)
	log.Println("c", keyBytes)
	if err != nil {
		panic(err)
	}
	return string(keyBytes)
}

func NewKeyboard(reader io.Reader) Keyboard {
	return &KeyboardImpl{reader: reader}
}
