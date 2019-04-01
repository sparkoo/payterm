package peripherals

import (
	"io"
	"log"
)

type Keyboard interface {
	Read() string
}

type KeyboardDummy struct {
	reader io.Reader
}

func (k *KeyboardDummy) Read() string {
	keyBytes := make([]byte, 1)
	_, err := k.reader.Read(keyBytes)
	log.Println("c", keyBytes)
	if err != nil {
		panic(err)
	}
	return string(keyBytes)
}

func NewDummyKeyboard(reader io.Reader) Keyboard {
	return &KeyboardDummy{reader: reader}
}
