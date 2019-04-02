package term

import "github.com/sparkoo/payterm/peripherals"

type termIO struct {
	display    *peripherals.Display
	buzzer     *peripherals.Buzzer
	keyboard   *peripherals.Keyboard
	cardReader *peripherals.CardReader
}
