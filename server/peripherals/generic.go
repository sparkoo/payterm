package peripherals

import (
	"fmt"
	"io"
	"log"
)

type peripheralWriter struct {
	writer io.Writer
}

func newWriter(writer io.Writer) *peripheralWriter {
	return &peripheralWriter{writer: writer}
}

func (w *peripheralWriter) write(message string) {
	if _, err := w.writer.Write([]byte(message)); err != nil {
		fmt.Printf("error writing [%v]", err)
	}
}

type peripheralReader struct {
	reader io.Reader
}

func newReader(reader io.Reader) *peripheralReader {
	return &peripheralReader{reader: reader}
}

func (r *peripheralReader) read() string {
	keyBytes := make([]byte, 64)

	_, err := r.reader.Read(keyBytes)
	log.Println("c", keyBytes)
	if err != nil {
		panic(err)
	}
	return string(keyBytes)
}
