package peripherals

import (
	"fmt"
	"io"
)

type OutputWriter interface {
	Write(string)
}

type OutputWriterImpl struct {
	writer io.Writer
}

func (ow *OutputWriterImpl) Write(message string) {
	if _, err := ow.writer.Write([]byte(message)); err != nil {
		fmt.Printf("error writing [%v]", err)
	}
}

func NewOutputWriter(writer io.Writer) OutputWriter {
	return &OutputWriterImpl{writer: writer}
}
