package peripherals

type CardReader interface {
	Read() string
}

type CardReaderImpl struct {
}

func (*CardReaderImpl) Read() string {
	return "dummy message 123"
}

func NewCardReader() CardReader {
	return &CardReaderImpl{}
}
