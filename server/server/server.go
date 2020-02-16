package server

import (
	"io"
)

type Server interface {
	AddWriteHandler(addr string) io.Writer
	AddReadListener(addr string) io.Reader
	AddHandler(addr string) io.ReadWriteCloser
	Start()
	Close()
}


