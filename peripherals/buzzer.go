package peripherals

import "fmt"

type Buzzer interface {
	Beep()
}

type BuzzerDummy struct {
}

func (*BuzzerDummy) Beep() {
	fmt.Println("buzzer beep !!!")
}

func NewDummyBuzzer() Buzzer {
	return &BuzzerDummy{}
}
