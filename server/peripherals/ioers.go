package peripherals

type InputReader interface {
	Read() string
}

type OutputWriter interface {
	Write(string)
}
