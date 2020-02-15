package server

import (
	"io"
)

type Server interface {
	AddHandler(addr string) io.ReadWriteCloser
	Start()
	Close()
}


