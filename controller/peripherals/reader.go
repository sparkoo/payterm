package peripherals

import (
	"io"
)

type InputReader interface {
	Read() string
}

type InputReaderImpl struct {
	reader io.Reader
}

func (ir *InputReaderImpl) Read() string {
	keyBytes := make([]byte, 64)

	_, err := ir.reader.Read(keyBytes)
	if err != nil {
		panic(err)
	}
	stringLength := 0
	for i := 0; i < 64; i++ {
		if keyBytes[i] == 0 {
			stringLength = i
			break
		}
	}
	//log.Println("c", keyBytes[:stringLength])
	return string(keyBytes[:stringLength])
}

func NewInputReader(reader io.Reader) InputReader {
	return &InputReaderImpl{reader: reader}
}
