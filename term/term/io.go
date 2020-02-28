package term

import "github.com/sparkoo/payterm/peripherals"

type termIO struct {
	display    peripherals.OutputWriter
	buzzer     peripherals.Buzzer
	keyboard   peripherals.InputReader
	cardReader peripherals.InputReader
}
