package peripherals

type Keyboard interface {
	Read() rune
}

type KeyboardDummy struct {
}

func (*KeyboardDummy) Read() rune {
	return '1'
}

func NewDummyKeyboard() Keyboard {
	return &KeyboardDummy{}
}
