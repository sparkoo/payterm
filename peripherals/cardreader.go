package peripherals

type CardReader interface {
	Read() string
}

type CardReaderDummy struct {
}

func (*CardReaderDummy) Read() string {
	return "dummy message 123"
}

func NewDummyCardReader() CardReader {
	return &CardReaderDummy{}
}
