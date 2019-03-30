package peripherals

import (
	"fmt"
	"io"
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

}